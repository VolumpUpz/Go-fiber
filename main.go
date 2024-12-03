package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var deptEmp []DepartmentEmployeeRequest

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.OriginalURL(), c.Method(), start)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

func main() {

	if err := godotenv.Load(); err != nil { //load env
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html") //set default path call template

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "benzlopster", Author: "Benz"})
	books = append(books, Book{ID: 2, Title: "benzlopster2", Author: "Benz2"})

	app.Post("/login", login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET"))},
	))
	app.Use(checkMiddleware) //route api ที่อยู่ต่อจาก app.Use จะถูกเรียกใช้ middleware

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

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{
	Email:    "user@example.com",
	Password: "password1234",
}

func login(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// if user.Email != memberUser.Email && user.Password != memberUser.Password {
	// 	return fiber.ErrUnauthorized
	// }

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims) //keep data of token and encrypt
	claims["email"] = user.Email
	claims["role"] = "admin"
	// if strings.Contains(user.Email, "ad") {
	// 	claims["role"] = "admin"
	// } else if strings.Contains(user.Email, "ur") {
	// 	claims["role"] = "user"
	// }

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}
