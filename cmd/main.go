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

	data, err := reader.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(data[1])

	// Initialize cassandra
	cassandra, err := cassandra.NewCassandra(os.Getenv("CASSANDRA_HOST"), os.Getenv("CASSANDRA_USERNAME"), os.Getenv("CASSANDRA_PASSWORD"))
	if err != nil {
		panic(err)
	}
	cassandra.Initialize()
}
