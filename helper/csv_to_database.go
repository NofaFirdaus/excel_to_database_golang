package helper

import (
	"encoding/csv"
	"exceltodatabases/config"
	"exceltodatabases/models"
	"fmt"
	"log"
	"os"
)

func CsvToDatabase(csvFileName string) error {
	// Membuka file CSV
	file, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal("Failed to open CSV file:", err)
	}
	fmt.Println("scv")
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV rows:", err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) != 6 {
			log.Printf("Skipping incomplete row on line %d: %v", i+1, row)
			continue
		}

		table := models.Table{
			Column1:   row[0],
			Column2:   row[1],
			Column3:   row[2],
			Column4:   row[3],
			Column5:   row[4],
			Column6:   row[5],
			Id_table1: 0,
		}

		queryTable1 := "INSERT INTO table1 (column1, column2, column3) VALUES (?, ?, ?)"
		table1, errTable1 := config.DB.Exec(queryTable1, table.Column1, table.Column2, table.Column3)
		if errTable1 != nil {
			log.Printf("Failed to insert row %v on line %d: %v\n", row, i+1, errTable1)
			continue
		}

		lastId, err := table1.LastInsertId()
		if err != nil {
			log.Printf("Failed to get LastInsertId for row %v on line %d: %v\n", row, i+1, err)
			continue
		}
		table.Id_table1 = int(lastId)

		queryTable2 := "INSERT INTO table2 (column1, column2, column3, id_table1) VALUES (?, ?, ?, ?)"
		_, errTable2 := config.DB.Exec(queryTable2, table.Column4, table.Column5, table.Column6, table.Id_table1)
		if errTable2 != nil {
			log.Printf("Failed to insert into table2 for row %v on line %d: %v\n", row, i+1, errTable2)
			continue
		}
	}

	defer func() {
		if err := os.Remove(csvFileName); err != nil {
			log.Printf("Failed to delete CSV file %s: %v\n", csvFileName, err)
		}
	}()

	return nil
}
