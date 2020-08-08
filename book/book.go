package book

import (
	"github/JamieBShaw/golang-fiber-rest-api/db"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := db.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)

}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {

	db := db.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(500).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn

	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found with id")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted")

}
