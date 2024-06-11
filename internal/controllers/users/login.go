package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	app_model "github.com/XanderMoroz/BookStore/internal/models"
	"github.com/XanderMoroz/BookStore/utils/token"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary        user authentication
// @Description    Authenticate User in app with given request body
// @Tags           Authentication
// @Accept         json
// @Produce        json
// @Param          request         	body        LoginInput    true    "Введите данные для авторизации"
// @Success        201              {string}    map[]
// @Failure        400              {string}    string    "Bad Request"
// @Router         /users/login 	[post]
func (h handler) Login(c *gin.Context) {

	log.Println("Поступил запрос на авторизацию в сервисе")
	var loginBody LoginInput

	if err := c.ShouldBindJSON(&loginBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		log.Println("Тело запроса успешно извлечено:")
		log.Printf("Username: <%s>\n Password: <%s>\n", loginBody.Username, loginBody.Password)
	}

	u := app_model.User{}

	u.Username = loginBody.Username
	u.Password = loginBody.Password

	token, err := h.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (h handler) LoginCheck(username string, password string) (string, error) {

	log.Println("Проверяем данные пользователя...")
	var err error

	u := app_model.User{}

	err = h.DB.Model(app_model.User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		log.Printf("... ошибка: <%v>", err)
		return "", err
	} else {
		log.Println("Пользователь - успешно извлечен:")
		log.Printf("Username: <%s>\n Password: <%s>\n", u.Username, u.Password)
	}

	err = VerifyPassword(password, u.Password)
	if err != nil {
		log.Printf("... ошибка: <%v>", err)
		return "", err
	}

	newToken, err := token.GenerateToken(u.ID)

	if err != nil {
		log.Printf("... ошибка: <%v>", err)
		return "", err
	}

	return newToken, nil

}

func VerifyPassword(password, hashedPassword string) error {

	log.Println("Верифицируем пароль...")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Invalid Password:", err)
		return err
	}
	log.Println("... Пароль успешно верифицирован")
	return nil

}
