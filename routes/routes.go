package routes

import (
	"github.com/mstomar698/go-auth/controllers"
	"github.com/gofiber/fiber"
)

func SetUp(app *fiber.App) {

	app.Get("/", controllers.Hello)
	app.Get("/", controllers.Test)

}
