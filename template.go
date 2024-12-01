package main

import "github.com/gofiber/fiber/v2"

func testHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, <b>World</b>!",
		"Name":  "Hello template html engine",
	})
}
