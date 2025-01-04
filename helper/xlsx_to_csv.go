package helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func XlsxToCsv(filePath string) (string, error) {
	fmt.Println(filePath)
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Failed to open Excel file:", err)
	}
	fmt.Println(filePath)

	sheetName := file.GetSheetName(0)
	fmt.Println(sheetName)
	rows, err := file.GetRows(sheetName)

	if err != nil {
		return "", fmt.Errorf("failed to create CSV file: %w", err)
	}

	csvFileName := "file/" + RandomString(5) + ".csv"
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		return "", fmt.Errorf("failed to create CSV file: %w", err)
	}

	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			log.Println("Failed to write row to CSV:", err)
		}
	}
	defer func() {
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("Failed to delete uploaded file %s: %v\n", filePath, err)
		}
	}()
	return csvFileName, nil
}
