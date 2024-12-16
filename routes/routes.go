package routes

import (
	"net/http"

	"go-api/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/user", handlers.CreateUser)
	http.HandleFunc("/user/update", handlers.UpdateUser)
	http.HandleFunc("/user/delete", handlers.DeleteUser)
}
