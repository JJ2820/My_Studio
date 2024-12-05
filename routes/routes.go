package routes

import (
	"net/http"
	"glofox-api/handlers"
)

func InitializeRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/classes", handlers.CreateClassHandler)
	router.HandleFunc("/bookings", handlers.BookClassHandler)
	return router
}
