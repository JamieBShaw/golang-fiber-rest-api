package main

import (
	"fmt"
	"github/JamieBShaw/golang-fiber-rest-api/book"
	"github/JamieBShaw/golang-fiber-rest-api/db"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloworld(c *fiber.Ctx) {

	c.Send("Hello, World!")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)

	app.Post("/api/v1/book/", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDB() {
	var err error
	db.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to db")
	}
	fmt.Println("Connected to database")

	db.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDB()
	defer db.DBConn.Close()

	setUpRoutes(app)
	app.Listen(3333)
}
