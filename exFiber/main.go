package main

import (
	"fmt"
	"log"

	"exFiber/managers/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	auth.Set(app)

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("[DEBUG] 2nd Handler")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		result := fmt.Sprintf("Hello, World! : %v", c.Locals("AID"))
		return c.SendString(result)
	})
	app.Post("/login", auth.Login)

	log.Fatal(app.Listen(":3000"))
}
