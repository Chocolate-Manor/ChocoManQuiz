package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", index)
	app.Get("/api/quizzes", getQuizzes)

	log.Fatal(app.Listen(":3000"))
}

func getQuizzes(c *fiber.Ctx) error {
	// a slice of maps with string as key and type of any
	list := []map[string]any{
		{
			"test": 123,
			"nope": true,
		},
		{
			"test": 12345,
		},
	}
	return c.JSON(list)
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
