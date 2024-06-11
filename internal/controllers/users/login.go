package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/XanderMoroz/BookStore/internal/models"
	"github.com/XanderMoroz/BookStore/utils"
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

	// Check if user exists
	var user models.User
	var defaulUserID uuid.UUID
	h.DB.Where("username = ?", loginBody.Username).First(&user)
	if user.ID == defaulUserID {
		log.Println("User not found")
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	log.Println("Верифицируем пароль...")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Failed to hash password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginBody.Password))
	if err != nil {
		log.Println("Invalid Password:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Invalid Password",
		})
		return
	}
	log.Println("... успешно")

	log.Println("Генерируем токен доступа...")

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Failed to generate token",
		})
		return
	} else {
		log.Println("... успешно:", token)
	}

	log.Println("Устанавливаем токен в куки пользователя...")
	c.SetCookie("jwt", token, 3600, "/", "127.0.0.1", false, true)

	log.Println("... успешно")

	// log.Println("Извлекаем токен из куков пользователя...")
	// cookie, err := c.Cookie("jwt")
	// if err != nil {
	// 	c.String(http.StatusNotFound, err.Error())
	// 	return
	// }
	// log.Println(cookie)

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "Token Created",
		"access_token": token,
	})

}

// func (h handler) LoginCheck(username string, password string) (string, error) {

// 	log.Println("Проверяем данные пользователя...")
// 	var err error

// 	u := models.User{}

// 	err = h.DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error

// 	if err != nil {
// 		log.Printf("... ошибка: <%v>", err)
// 		return "", err
// 	} else {
// 		log.Println("Пользователь - успешно извлечен:")
// 		log.Printf("Username: <%s>\n Password: <%s>\n", u.Username, u.Password)
// 	}

// 	// err = VerifyPassword(password, u.Password)
// 	// if err != nil {
// 	// 	log.Printf("... ошибка: <%v>", err)
// 	// 	return "", err
// 	// }

// 	// newToken, err := utils.GenerateToken(u.ID)

// 	// if err != nil {
// 	// 	log.Printf("... ошибка: <%v>", err)
// 	// 	return "", err
// 	// }

// 	// return newToken, nil

// }

// func VerifyPassword(password, hashedPassword string) error {

// 	log.Println("Верифицируем пароль...")
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	if err != nil {
// 		log.Println("Invalid Password:", err)
// 		return err
// 	}
// 	log.Println("... Пароль успешно верифицирован")
// 	return nil

// }
