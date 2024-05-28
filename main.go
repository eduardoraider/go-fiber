package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://wookye.com.br, http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	v1 := app.Group("/api/v1")
	v2 := app.Group("/v2")
	user := v2.Group("/user")
	user.Use(func(c *fiber.Ctx) error {
		c.Set("X-Frame-Options", "SAMEORIGIN")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!\n")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong\n")
	})

	app.Get("/names.txt", func(c *fiber.Ctx) error {
		return c.SendString("John, Mary, Bob, Chloe\n")
	})

	user.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("resource user\n")
	})

	app.Get("/user/*", func(c *fiber.Ctx) error {
		s := c.Params("*")
		if s == "" {
			s = "none!"
		}
		fmt.Fprintf(c, "%s\n", s)
		return nil
	})

	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s-%s\n", c.Params("from"), c.Params("to"))
		return nil
	})

	app.Get("/plants/:type.:specie", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s.%s\n", c.Params("type"), c.Params("specie"))
		return nil
	})

	v1.Get("/hello/:name?", func(c *fiber.Ctx) error {
		s := c.Params("name")
		if s == "" {
			s = "Hello"
		}
		fmt.Fprintf(c, "Hello, %s\n", s)
		return nil
	})

	app.Listen(":3000")
}
