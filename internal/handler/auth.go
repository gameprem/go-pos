package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-pos/internal/database"
	"go-pos/internal/model"
)

var jwtSecret = []byte("supersecretkey")

// Register godoc
// @Summary Register user
// @Description Register a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user  body  model.User  true  "User info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username and password required"})
	}
	if err := database.DB.Create(user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username already exists"})
	}
	return c.JSON(fiber.Map{"message": "Registration successful"})
}

// Login godoc
// @Summary Login user
// @Description Login and get JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user  body  model.User  true  "User info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	input := new(model.User)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	var user model.User
	if err := database.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}
	return c.JSON(fiber.Map{"token": signedToken})
}
