package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not found item")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	c.BodyParser(book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)

}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)

	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			//books ที่ตำแหน่งเริ่มต้นจนถึง ตำแหน่งก่อนสิ้นสุด , books ที่ตำแหน่งของ i +1 ไปจนตำแหน่งสุดท้าย เช่น [1,2,3,4,5] [1,2] + [4,5]
			//append ไม่ support slice ด้วยกัน จึงต้องกระจาย slice ออกเป็น element แต่ละตัว โดยการใช้ ...
			books = append(books[:i], books[i+1:]...)
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
