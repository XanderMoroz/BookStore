package users

import (
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// Определяем структуру тела запроса на создание пользователя
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary			user registration
// @Description		Register User in app with given request body
// @Tags			Authentication
// @Accept			json
// @Produce			json
// @Param			request				body		RegisterInput	true	"Введите данные для регистрации"
// @Success			201					{string}	map[string]string
// @Failure			400					{string}	string	"Bad Request"
// @Router			/users/register 	[post]
func (h handler) Register(c *gin.Context) {

	log.Println("Поступил запрс на регистрацию в сервисе")
	request := RegisterInput{}

	// Пытаемся получить тело запроса
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"body parsing error": err.Error()})
		return
	} else {
		log.Println("Тело запроса успешно извлечено:")
		log.Printf("Username: <%v>\n Password: <%v>\n", request.Username, request.Password)
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
	} else {
		log.Printf("Пароль успешно зашифрован: <%v>\n", hashedPassword)
	}
	u.Password = string(hashedPassword)

	log.Println("Добавляем нового пользователя в БД...")
	newUser := h.DB.Create(&u)
	if newUser.Error != nil {
		// Обрабатываем ошибку
		log.Fatal("Пользователя создать не удалось", newUser.Error)
		c.AbortWithError(http.StatusNotFound, newUser.Error)
		return
	}

	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	c.JSON(http.StatusCreated, &newUser)
}
