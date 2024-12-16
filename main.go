package main

import (
	"fmt"
	"go-api/database"
	"go-api/routes"
	"net/http"
)

func main() {
	database.Connect()

	routes.SetupRoutes()

	fmt.Println("Server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}
