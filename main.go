package main

import (
	"log"

	db "banc-api/src/infrastructure/db/adapter"

	"github.com/gofiber/fiber/v3"
)

func main() {

	db.DBconection()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hello world")
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
