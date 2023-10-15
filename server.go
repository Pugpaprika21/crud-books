package main

import (
	booksController "go-basic/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	engine.AddFunc("RowsNumber", func(i int) int {
		return i + 1
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	/* show all books*/
	app.Get("/books-list", booksController.GetBooks)

	app.Listen(":3000")
}
