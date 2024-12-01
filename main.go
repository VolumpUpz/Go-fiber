package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var deptEmp []DepartmentEmployeeRequest

var books []Book

func main() {

	app := fiber.New()

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

	app.Listen(":8080")
}
