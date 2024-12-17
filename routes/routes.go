package routes

import (
	"net/http"

	"go-api/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/users", handlers.GetAll)
	http.HandleFunc("/user", handlers.Create)
	http.HandleFunc("/user/update", handlers.Update)
	http.HandleFunc("/user/delete", handlers.Delete)
}
