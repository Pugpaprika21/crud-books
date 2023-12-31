package booksController

import (
	database "go-basic/helpers"
	BooksModel "go-basic/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

var db, _ = database.DBConnection()

func GetBooks(c *fiber.Ctx) error {
	res, err := db.Query("select * from books")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil
	}
	defer res.Close()

	var booksList []BooksModel.Books
	for res.Next() {
		var books BooksModel.Books

		err := res.Scan(&books.Id, &books.Title, &books.Author)
		if err != nil {
			log.Fatalf("Error scanning rows: %v", err)
			return nil
		}

		booksList = append(booksList, books)
	}

	return c.Render("index", fiber.Map{
		"Books": booksList,
	})
}
