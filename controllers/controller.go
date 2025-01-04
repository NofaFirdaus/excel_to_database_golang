package controllers

import (
	"encoding/json"
	"exceltodatabases/helper"
	"exceltodatabases/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dstPath := "file/" + handler.Filename
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// fmt the size of the created file
	fileInfo, err := os.Stat(dstPath)
	if err != nil {
		fmt.Printf("Error getting file info: %v", err)
	} else {
		fmt.Printf("File created: %s, Size: %d bytes", dstPath, fileInfo.Size())
	}

	// Attempt to convert the Excel file to CSV
	csvFileName, err := helper.XlsxToCsv(dstPath)
	if err != nil {
		fmt.Printf("Failed to convert Excel to CSV: %v", err)
		http.Error(w, "Failed to convert Excel file", http.StatusInternalServerError)
		return
	}

	// Process the CSV file
	errDa := helper.CsvToDatabase(csvFileName)
	if errDa != nil {
		fmt.Printf("Failed to process CSV file: %v", errDa)
		http.Error(w, "Failed to process CSV file", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetAllTable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	table := models.GetAll()
	// response := map[string]any{"tabel": table}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(table)

}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	models.DeleteAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("oke")
}
