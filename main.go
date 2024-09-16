package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("home page")
	})

	var arr [3]int

	arr[0] = 3
	arr[4] = 2

	app.Listen("localhost:8080")

}
