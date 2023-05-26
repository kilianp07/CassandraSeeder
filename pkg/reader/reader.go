package reader

import (
	"encoding/csv"
	"os"

	"github.com/google/uuid"
	"github.com/kilianp07/CassandraSeeder/utils/structs"
)

func Read(filepath string) ([]structs.Contact, error) {
	var data []structs.Contact
	// Read csv file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range csvLines {
		data = append(data, structs.Contact{
			Id:          uuid.New().String(),
			Title:       line[0],
			Name:        line[1],
			Address:     line[2],
			RealAddress: line[3],
			Departement: line[4],
			Country:     line[5],
			Tel:         line[6],
			Email:       line[7],
		})
	}
	return data, nil
}
