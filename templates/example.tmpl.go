package templates

var MAIN_TEMPLATE string = `package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
`

var EXAMPLE_MIDDLEWARE_TEMPLATE string = `package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func HeartBeat(endpoint string) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		methodName := ctx.Method()
		path := ctx.Path()

		if (methodName == "GET" || methodName == "HEAD") && strings.EqualFold(path, endpoint) {
			return ctx.Status(http.StatusOK).SendString(".")
		}

		return ctx.Next()
	}
}
`
