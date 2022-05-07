package main

import (
	"export/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/ping", handler.Ping)

	_ = app.Listen(":5000")
}
