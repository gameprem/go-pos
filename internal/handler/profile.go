package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-pos/internal/database"
	"go-pos/internal/model"
)

// Me godoc
// @Summary Get current user profile
// @Description Get profile of the authenticated user
// @Tags profile
// @Security bearerAuth
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/profile/me [get]
func Me(c *fiber.Ctx) error {
	claims := c.Locals("user")
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}
	userIDFloat, ok := mapClaims["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user id in token"})
	}
	userID := uint(userIDFloat)

	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	var userInfo model.UserInfo
	database.DB.Where("user_id = ?", userID).First(&userInfo)

	return c.JSON(fiber.Map{
		"id":         user.ID,
		"username":   user.Username,
		"first_name": userInfo.FirstName,
		"last_name":  userInfo.LastName,
	})
}
