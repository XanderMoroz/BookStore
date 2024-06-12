package users

import (
	"log"
	"time"

	"github.com/gofiber/fiber"
)

// @Summary		logout current user
// @Description Clear JWT token by setting an empty value and expired time in the cookie
// @Tags 		Authentication
// @ID			logout-current-user
// @Produce		json
// @Success		200		{string}	map[]
// @Router		/api/v1/logout [get]
func Logout(c *fiber.Ctx) error {
	log.Println("Получен запрос на выход из аккаунта")

	log.Println("Удаляем токен из куков...")
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expired 1 hour ago
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)
	log.Println("... успешно")

	// Return success response indicating logout was successful
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
