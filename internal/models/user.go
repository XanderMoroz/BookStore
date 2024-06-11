package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/XanderMoroz/BookStore/utils"
)

var DB *gorm.DB

// type User struct {
// 	gorm.Model
// 	Username string `gorm:"size:255;not null;unique" json:"username"`
// 	Password string `gorm:"size:255;not null;" json:"password"`
// }

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primary_key"` //`gorm:"type:uuid"`
	Name      string    `gorm:"not null"`
	Username  string    `gorm:"type:char(128);uniqueIndex"`
	Password  []byte    `json:"-"` // contain the hashed password.
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary		get current user
// @Description Get token from users cookee
// @Tags 		Authentication
// @ID			get-current-user
// @Produce		json
// @Success		200		{object}	UserResponse
// @Router		/api/admin/user [get]
func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func GetUserByID(uid uint) (User, error) {

	u := User{}

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	// u.PrepareGive()

	return u, nil

}

func (u User) PrepareGive() {

}
