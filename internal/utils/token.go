package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/XanderMoroz/BookStore/config"
)

func GenerateToken(userID uuid.UUID) (string, error) {

	env := config.NewEnv()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})

	secretKey := []byte(env.AccessTokenSecret)
	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

// func GenerateToken(user_id uint) (string, error) {

// 	log.Println("Начинаем генерацию токена доступа...")

// 	env := config.NewEnv()

// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["user_id"] = user_id
// 	claims["exp"] = time.Now().Add(time.Hour * time.Duration(env.AccessTokenExpiryHour)).Unix()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return token.SignedString([]byte(env.AccessTokenSecret))

// }

func TokenValid(c *gin.Context) error {

	log.Println("Валидируем токен...")
	jwtFromCookie, err := c.Cookie("jwt")
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		// return jwtFromCookie
	}
	env := config.NewEnv()
	hmacSecret := []byte(env.AccessTokenSecret)

	token, err := jwt.Parse(jwtFromCookie, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена и другие параметры
		return hmacSecret, nil
	})

	if err != nil {
		// log.Printf("При извлечении токена произошла ошибка <%v>\n", err)
		return fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Неверный JWT токен")
		return fmt.Errorf("invalid JWT Token ")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		log.Println("Не удалось извлечь USER_ID из токена")
		return fmt.Errorf("failed to parse claims")
	} else {
		log.Printf("USER_ID: <%s>\n", userID)
	}
	return nil
}

func ExtractUserIDFromToken(c *gin.Context) (string, error) {

	jwtFromCookie, err := c.Cookie("jwt")
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		// return jwtFromCookie
	}
	env := config.NewEnv()
	hmacSecret := []byte(env.AccessTokenSecret)

	token, err := jwt.Parse(jwtFromCookie, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена и другие параметры
		return hmacSecret, nil
	})

	if err != nil {
		// log.Printf("При извлечении токена произошла ошибка <%v>\n", err)
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Неверный JWT токен")
		return "", fmt.Errorf("invalid JWT Token ")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		log.Println("Не удалось извлечь USER_ID из токена")
		return "", fmt.Errorf("failed to parse claims")
	} else {
		log.Printf("USER_ID: <%s>\n", userID)
		return userID, nil
	}

}
