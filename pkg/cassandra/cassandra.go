package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
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

	c.CreateKeyspace()
	c.cluster.Keyspace = c.keyspace

	session, err := c.cluster.CreateSession()

	defer session.Close()

	if err != nil && session != nil {
		fmt.Printf("Error creating session: %v \n", err)
		return err
	}
	fmt.Println("cassandra well initialized")
	return nil
}

// Create keyspace
func (c *Cassandra) CreateKeyspace() error {
	session, err := c.cluster.CreateSession()
	defer session.Close()

	err = session.Query(`
		CREATE KEYSPACE IF NOT EXISTS ` + c.keyspace + `
		WITH REPLICATION = {
			'class' : 'SimpleStrategy',
			'replication_factor' : 1
		}
	`).Exec()
	if err != nil {
		fmt.Println("Error creating keyspace: ", err)
		return err
	}
	fmt.Println("Keyspace is created successfully.")
	return nil
}

func (c *Cassandra) Migrate() error {
	session, err := c.cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return err
	}
	defer session.Close()

	createTableQuery := `CREATE TABLE IF NOT EXISTS my_table (
		id UUID PRIMARY KEY,
		title text,
		name text,
		address text,
		realAddress text,
		department text,
		country text,
		tel text,
		email text
	)`

	err = session.Query(createTableQuery).Exec()
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	fmt.Println("Table created successfully")

	return nil
}

func (c *Cassandra) MigrateData(data structs.Contact) error {

	session, err := c.cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return err
	}
	defer session.Close()

	// Insert restaurant data
	err = session.Query(`
		INSERT INTO my_table (id, title, name, address, realAddress, department, country, tel, email)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		data.Id, data.Title, data.Name, data.Address, data.RealAddress, data.Departement, data.Country, data.Tel, data.Email,
	).Exec()
	if err != nil {
		return err
	}
	return nil
}
