package models

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/XanderMoroz/BookStore/utils/token"
)

var DB *gorm.DB

// type User struct {
// 	gorm.Model
// 	ID        uuid.UUID `gorm:"type:uuid"`
// 	Username  string    `gorm:"not null"`
// 	Email     string    `gorm:"uniqueIndex"`
// 	Password  []byte    `json:"-"` // contain the hashed password.
// 	CreatedAt time.Time `gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime"`
// 	// DeletedAt gorm.DeletedAt
// 	// Articles  []Article `gorm:"foreignKey:UserID"`
// }

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u User) PrepareGive() {

}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

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

	u.PrepareGive()

	return u, nil

}
