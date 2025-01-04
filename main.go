package main

import (
	"exceltodatabases/config"
	"exceltodatabases/controllers"
	"net/http"
)

func main() {
	config.Database()

	http.HandleFunc("/", controllers.Welcome)
	http.HandleFunc("/add", controllers.Add)
	http.HandleFunc("/get", controllers.GetAllTable)
	http.ListenAndServe(":8080", nil)
}
