package main

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  BearerAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

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

	fmt.Println("Starting web server on port :3030")
	fmt.Println("Docs : http://localhost:3030/docs")

	log.Fatal(app.Listen(":3030"))
}
