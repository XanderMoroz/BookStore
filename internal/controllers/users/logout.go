package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		logout current user
// @Description Clear JWT token by setting an empty value and expired time in the cookie
// @Tags 		Authentication
// @ID			logout-current-user
// @Produce		json
// @Success		200		{string}	map[]
// @Router		/users/logout [get]
func (h handler) Logout(c *gin.Context) {
	log.Println("Получен запрос на выход из аккаунта")

	log.Println("Удаляем токен из куки пользователя...")
	c.SetCookie("jwt", "", 3600, "/", "127.0.0.1", false, true)
	log.Println("... успешно")

	// Return success response indicating logout was successful
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout successful",
	})

	return
}
