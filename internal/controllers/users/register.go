package users

import (
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// Определяем структуру тела запроса на создание пользователя
type SignUpUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// @Summary			user registration
// @Description		Register User in app with given request body
// @Tags			Authentication
// @Accept			json
// @Produce			json
// @Param			request				body		SignUpUserRequest	true	"Введите данные для регистрации"
// @Success			201					{string}	map[string]string
// @Failure			400					{string}	string	"Bad Request"
// @Router			/users/register 	[post]
func (h handler) Register(c *gin.Context) {

	log.Println("Поступил запрс на регистрацию в сервисе")

	request := SignUpUserRequest{}

	// Пытаемся получить тело запроса
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"body parsing error": err.Error()})
		return
	} else {
		log.Println("Тело запроса успешно извлечено:")
		log.Printf("Username: <%v>\n Password: <%v>\n", request.Username, request.Password)
	}

	// Создаем новый экземпляр пользователя
	newUser := models.User{}

	// Чистим и присваиваем ему значения из тела запроса
	newUser.Username = html.EscapeString(strings.TrimSpace(request.Username))

	// Шифруем пароль из запроса
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка шифрования пароля")
		return
	} else {
		log.Printf("Пароль успешно зашифрован: <%v>\n", hashedPassword)
	}

	// Присваиваем новому пользователю уникальный ID
	newUser.ID = uuid.New()
	newUser.Name = request.Name
	newUser.Password = hashedPassword

	log.Printf("Добавляем в БД нового пользователя:")
	log.Printf("	ID: <%+v>\n", newUser.ID)
	log.Printf("	Имя: <%+v>\n", newUser.Name)
	log.Printf("	Username: <%+v>\n", newUser.Username)
	log.Printf("	hashedPassword: <%s>\n", string(hashedPassword))

	result := h.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error.Error())
		return
	} else {
		log.Println("Новый пользователь успешно зарегистрирован")
		log.Printf("	ID: <%v>\n", newUser.ID)
		log.Printf("	Имя: <%+v>\n", newUser.Name)
		log.Printf("	Username: <%s>\n", newUser.Username)
		log.Printf("	Хэш-пароль: <%s>\n", string(newUser.Password))
	}

	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"new_user": &newUser,
	})

}
