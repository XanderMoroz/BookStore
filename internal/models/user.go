package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
