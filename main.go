package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kilianp07/CassandraSeeder/pkg/cassandra"
	"github.com/kilianp07/CassandraSeeder/pkg/reader"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	data, err := reader.Read("./contacts.csv")
	if err != nil {
		panic(err)
	}

	// Initialize cassandra
	cassandra, err := cassandra.NewCassandra(os.Getenv("CASSANDRA_HOST"), os.Getenv("CASSANDRA_USERNAME"), os.Getenv("CASSANDRA_PASSWORD"), "contacts")
	if err != nil {
		panic(err)
	}
	if err := cassandra.Initialize(); err != nil {
		panic(err)
	}

	if err := cassandra.Migrate(); err != nil {
		panic(err)
	}

	for _, contact := range data {
		fmt.Println("Migrating Contact: ", contact.Name)
		if err := cassandra.MigrateData(contact); err != nil {
			panic(err)
		}
	}
}
