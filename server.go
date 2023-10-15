package main

import (
	booksController "go-basic/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go Fiber")
	})

	app.Get("/get-books", booksController.GetBooks)

	app.Listen(":3000")
}
