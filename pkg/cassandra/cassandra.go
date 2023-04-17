package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/kilianp07/CassandraSeeder/utils/structs"
)

type Cassandra struct {
	cluster *gocql.ClusterConfig

	host     string
	username string
	password string
	keyspace string
}

// NewCassandra creates a new Cassandra object with the given host, username, password and keyspace.
func NewCassandra(host string, username string, password string, keyspace string) (*Cassandra, error) {
	return &Cassandra{
		host:     host,
		username: username,
		password: password,
		keyspace: keyspace,
	}, nil
}

func (c *Cassandra) Initialize() error {
	c.cluster = gocql.NewCluster(c.host)
	c.cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.username,
		Password: c.password,
	}

	c.cluster.Keyspace = c.keyspace

	session, err := c.cluster.CreateSession()
	if err != nil && session != nil {
		fmt.Printf("Error creating session: %v \n", err)
		return err
	}
	fmt.Println("cassandra well initialized")
	return nil
}

func (c *Cassandra) Migrate() error {
	session, err := c.cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return err
	}
	defer session.Close()

	// Create tables if not exists
	err = session.Query(`
		CREATE TABLE IF NOT EXISTS restaurants (
			restaurant_id TEXT PRIMARY KEY,
			name TEXT,
			borough TEXT,
			cuisine TEXT,
		)
	`).Exec()
	if err != nil {
		fmt.Println("Error creating table restaurants: ", err)
		return err
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS addresses (
			address_id TEXT PRIMARY KEY,
			restaurant_id TEXT,
			building TEXT,
			street TEXT,
			zipcode TEXT
		)
	`).Exec()
	if err != nil {
		fmt.Println("Error creating table addresses: ", err)
		return err
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS grades (
			grade_id TEXT PRIMARY KEY,
			restaurant_id TEXT,
			date TIMESTAMP,
			grade TEXT,
			score INT
		)
	`).Exec()
	if err != nil {
		fmt.Println("Error creating table grades: ", err)
		return err
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS coordinates (
			coord_id TEXT PRIMARY KEY,
			address_id TEXT,
			type TEXT,
			coordinates LIST<FLOAT>
		)
	`).Exec()
	if err != nil {
		fmt.Println("Error creating table coordinates: ", err)
		return err
	}

	fmt.Println("Tables are created successfully.")

	return nil
}

func (c *Cassandra) MigrateRestaurantData(r structs.Restaurant) error {

	session, err := c.cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return err
	}
	defer session.Close()

	// Generate UUIDs for primary keys
	addressID := uuid.New().String()
	gradeID := uuid.New().String()
	coordID := uuid.New().String()

	// Insert restaurant data
	err = session.Query(`
		INSERT INTO restaurants (restaurant_id, name, borough, cuisine)
		VALUES (?, ?, ?, ?)`,
		r.RestaurantID, r.Name, r.Borough, r.Cuisine,
	).Exec()
	if err != nil {
		return err
	}

	// Insert address data
	err = session.Query(`
		INSERT INTO addresses (address_id, restaurant_id, building, street, zipcode)
		VALUES (?, ?, ?, ?, ?)`,
		addressID, r.RestaurantID, r.Address.Building, r.Address.Street, r.Address.Zipcode,
	).Exec()
	if err != nil {
		return err
	}

	// Insert grade data
	for _, g := range r.Grades {
		err = session.Query(`
			INSERT INTO grades (grade_id, restaurant_id, date, grade, score)
			VALUES (?, ?, ?, ?, ?)`,
			gradeID, r.RestaurantID, g.Date.Date, g.Grade, g.Score,
		).Exec()
		if err != nil {
			return err
		}

		// Generate a new UUID for the next grade
		gradeID = uuid.New().String()
	}

	// Insert coordinates data
	err = session.Query(`
		INSERT INTO coordinates (coord_id, address_id, type, coordinates)
		VALUES (?, ?, ?, ?)`,
		coordID, addressID, r.Address.Coord.Type, r.Address.Coord.Coordinates,
	).Exec()
	if err != nil {
		return err
	}

	return nil
}
