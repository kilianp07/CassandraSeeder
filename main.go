package main

import (
	"os"

	"github.com/kilianp07/CassandraSeeder/pkg/cassandra"
	"github.com/kilianp07/CassandraSeeder/pkg/reader"
)

func main() {

	data, err := reader.Read(os.Getenv("JSON_FILEPATH"))
	if err != nil {
		panic(err)
	}

	// Initialize cassandra
	cassandra, err := cassandra.NewCassandra(os.Getenv("CASSANDRA_HOST"), os.Getenv("CASSANDRA_USERNAME"), os.Getenv("CASSANDRA_PASSWORD"), os.Getenv("CASSANDRA_KEYSPACE"))
	if err != nil {
		panic(err)
	}
	if err := cassandra.Initialize(); err != nil {
		panic(err)
	}

	if err := cassandra.Migrate(); err != nil {
		panic(err)
	}

	for _, restaurant := range data {
		if err := cassandra.MigrateRestaurantData(restaurant); err != nil {
			panic(err)
		}
	}
}
