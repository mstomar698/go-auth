package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
)

func main() {
	// for db
	database.Connect()
	// h := handlers.New(DB)

	// for fiber
	app := fiber.New()

	// for routes
	routes.SetUp(app)

	app.Listen(":8000")
}
