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
}

func NewCassandra(host string, username string, password string) (*Cassandra, error) {

	return &Cassandra{
		host:     host,
		username: username,
		password: password,
	}, nil
}

func (c *Cassandra) Initialize() error {
	c.cluster = gocql.NewCluster(c.host)
	c.cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.username,
		Password: c.password,
	}

	c.cluster.Keyspace = c.cluster.Keyspace

	session, err := c.cluster.CreateSession()
	if err != nil {
		fmt.Printf("Error creating session: %v \n", err)
		return err
	}
	fmt.Println("cassandra well initialized", session)
	return nil
}
