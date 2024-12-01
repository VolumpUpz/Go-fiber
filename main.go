package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var deptEmp []DepartmentEmployeeRequest

var books []Book

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "benzlopster", Author: "Benz"})
	books = append(books, Book{ID: 2, Title: "benzlopster2", Author: "Benz2"})

	app.Get("/books", getBooks)
	app.Get("books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	deptEmp = append(deptEmp, DepartmentEmployeeRequest{
		DepartmentDetail: Department{ID: 1, Name: "Developer"},
		EmployeeDetail:   Employee{ID: 1, Name: "Benz", Salary: 1000000},
	})

	// app.Get("/deptemps", getDeptEmp)

	app.Get("/deptemps", getDeptEmp)

	app.Post("/upload", uploadFile)

	app.Get("/test-html", testHtml)

	app.Get("/config", getEnv)

	app.Listen(":8080")
}

func getEnv(c *fiber.Ctx) error {
	if secret := os.Getenv("SECRET"); secret == "" {
		c.JSON(fiber.Map{
			"SECRET": "default secret",
		})
	}

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}
