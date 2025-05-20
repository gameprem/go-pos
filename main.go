package main

import (
	"fmt"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-pos/internal/database"
	"go-pos/internal/handler"
	"go-pos/internal/middleware"
	"log"
)

func main() {
	database.Init()
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	app.Get("/protected", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		claims := c.Locals("user")
		return c.JSON(fiber.Map{"message": "You are authenticated!", "claims": claims})
	})

	app.Get("/docs", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Type("html").SendString(htmlContent)
	})

	fmt.Printf("Starting web server on port :3030")

	log.Fatal(app.Listen(":3030"))
}
