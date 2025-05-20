package router

import (
	"github.com/gofiber/fiber/v2"
	"go-pos/internal/handler"
	"go-pos/internal/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

	profile := v1.Group("/profile")
	profile.Get("/me", middleware.JWTMiddleware, handler.Me)
	// เพิ่ม group อื่นๆ ได้ที่นี่ เช่น v1.Group("/product")
}
