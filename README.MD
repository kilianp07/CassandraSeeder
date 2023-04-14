# CassandraSeeder

*CassandraSeeder* is a Go application that demonstrates data migration to Cassandra using the gocql driver. It provides a simple example of how to migrate data from a Go struct to Cassandra tables.

## Prerequisites

Before running the *CassandraSeeder* application, make sure you have the following installed:

- Go (1.15 or later)
- Cassandra (3.0 or later)
- Run validate_json script to validate the json file
```sh
bash validate_json.sh
```
## Getting Started

1. Clone the repository to your local machine:
```sh
git clone https://github.com/kilianp07/CassandraSeeder.git
```
2. Navigate to the project directory:
```sh
cd CassandraSeeder
```

3. Build the application:
```sh
make build
```

4. Create your .env file structured like this:
```sh
CASSANDRA_HOST=
CASSANDRA_KEYSPACE=
CASSANDRA_USERNAME=
CASSANDRA_KEYSPACE =
JSON_FILEPATH=
```
Note: The .env file must be in the same directory as the executable. 


5. Run the application:
```sh
./build/CassandraSeeder
```


