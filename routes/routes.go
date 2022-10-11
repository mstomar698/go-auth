package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func SetUp(app *fiber.App) {

	app.Get("/", controllers.Hello)

}
