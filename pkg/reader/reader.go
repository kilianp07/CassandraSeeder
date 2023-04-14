package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kilianp07/CassandraSeeder/utils/structs"
)

func Read(filepath string) ([]structs.Restaurant, error) {
	// Open the JSON file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return nil, err
	}
	defer file.Close()

	// Read the JSON data from the file
	jsonData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return nil, err
	}

	// Unmarshal JSON data into slice of Restaurant structs
	var restaurants []structs.Restaurant
	err = json.Unmarshal(jsonData, &restaurants)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return restaurants, nil
}
