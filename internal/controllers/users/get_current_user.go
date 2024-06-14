package users

import (
	"net/http"

	"github.com/XanderMoroz/BookStore/internal/utils"
	"github.com/gin-gonic/gin"
)

// @Summary		get current user
// @Description Get token from users cookee
// @Tags 		Authentication
// @ID			get-current-user
// @Produce		json
// @Success		200		{object}	models.UserResponse
// @Router		/api/admin/user [get]
func (h handler) CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractUserIDFromToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := utils.GetUserByIDFromDB(user_id)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": currentUser})
}
