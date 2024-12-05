package main

import (
	"fmt"
	"net/http"
	"glofox-api/routes"
)

func main() {
	router := routes.InitializeRoutes()

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
