package main

import (
	"github.com/mstomar698/go-auth/database"
	"github.com/mstomar698/go-auth/routes"
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
