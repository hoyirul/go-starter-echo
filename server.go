package main

import (
	"go-echo/db"
	"go-echo/routes"
)

func main() {
	// Initialize the database
	db.Init()

	// Initialize routes
	e := routes.Init()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
