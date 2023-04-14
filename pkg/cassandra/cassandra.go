package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
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
	if err != nil {
		fmt.Printf("Error creating session: %v \n", err)
		return err
	}
	fmt.Println("cassandra well initialized", session)
	return nil
}

func (c *Cassandra) Migrate() error {
	// Connect to Cassandra
	cluster := gocql.NewCluster(c.host)
	cluster.Keyspace = c.keyspace
	session, err := cluster.CreateSession()
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
			address_id TEXT
		)
	`).Exec()
	if err != nil {
		fmt.Println("Error creating table restaurants: ", err)
		return err
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS addresses (
			address_id TEXT PRIMARY KEY,
			building TEXT,
			coord_type TEXT,
			coord_coords LIST<FLOAT>,
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
			grade_id UUID PRIMARY KEY,
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
			coord_id UUID PRIMARY KEY,
			type TEXT,
			coordinates LIST<FLOAT>
		)
	`).Exec()
	if err != nil {
		fmt.Println("Error creating table coordinates: ", err)
		return err
	}

	fmt.Println("Migration completed successfully.")

	return nil
}
