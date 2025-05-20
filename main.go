package main

// @title Simple API
// @version 1.0
// @description Example API
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization

import (
	"fmt"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-pos/internal/database"
	"go-pos/internal/router"
	"log"
)

func main() {
	database.Init()
	app := fiber.New()
	app.Use(logger.New())

	router.Setup(app)

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
