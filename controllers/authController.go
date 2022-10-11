package controllers

import "github.com/gofiber/fiber"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World! This is a test.")
}

// create a fiber function using the fiber context
func Test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name": "John",
		"age":  33,
	})
}
