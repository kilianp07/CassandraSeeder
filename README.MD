# CassandraSeeder

*CassandraSeeder* is a Go application that demonstrates data migration to Cassandra using the gocql driver. It provides a simple example of how to migrate data from a Go struct to Cassandra tables.

## Downloading a Release

You can download a release from the [releases page](https://github.com/kilianp07/CassandraSeeder/releases). 


## Prerequisites (Building from Source)

Before running the *CassandraSeeder* application, make sure you have the following installed:

- Go (1.15 or later)
- Cassandra (3.0 or later)


## Getting Started

1. Using a release:
    * Download the latest executable from the [releases page](https://github.com/kilianp07/CassandraSeeder/releases)
    * Create a `.env` file structured like this:
    ```sh
    CASSANDRA_HOST=
    CASSANDRA_USERNAME=
    CASSANDRA_PASSWORD=
    ```
    * Place the `.env` and the `contact.csv` files in the same directory as the executable.
    * Go to your directory and run the executable:
    ```sh
    ./CassandraSeeder
    ```


2. Building from source:
    * Clone the repository:
    ```sh
    git clone https://github.com/kilianp07/CassandraSeeder.git
    ```
    * Navigate to the project directory:
    ```sh
    cd CassandraSeeder
    ```

    * Build the application:
    ```sh
    make build
    ```

    *. Create your `.env` file structured like this:
    ```sh
    CASSANDRA_HOST=
    CASSANDRA_USERNAME=
    CASSANDRA_PASSWORD=
    ```
    Note: The `.env` and `contact.csv` files must be in the same directory as the executable. 


    * Run the application:
    ```sh
    ./build/CassandraSeeder
    ```


