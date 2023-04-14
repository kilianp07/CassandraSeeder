package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kilianp07/CassandraSeeder/pkg/cassandra"
	"github.com/kilianp07/CassandraSeeder/pkg/reader"
)

func main() {

	// Read .env file
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getenv("CASSANDRA_HOST"))

	data, err := reader.Read()
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
}
