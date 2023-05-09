package users

import (
	"html"
	"log"
	"net/http"
	"strings"

	"BookStore/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Определяем структуру тела запроса на создание пользователя
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Register(c *gin.Context) {

	request := RegisterInput{}

	// Пытаемся получить тело запроса
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Создаем новый экземпляр пользователя
	u := models.User{}

	// Чистим и присваиваем ему значения из тела запроса
	u.Username = html.EscapeString(strings.TrimSpace(request.Username))

	// Шифруем пароль из запроса
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка шифрования пароля")
		return
	}
	u.Password = string(hashedPassword)

	// Пытаемся создать нового пользователя
	newUser := h.DB.Create(&u)

	// Обрабатываем ошибку
	if newUser.Error != nil {
		c.AbortWithError(http.StatusNotFound, newUser.Error)
		return
	}

	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	c.JSON(http.StatusCreated, &newUser)

}
