package main

import (
	"log"
	"net/http"

	"github.com/XanderMoroz/BookStore/config"
	"github.com/XanderMoroz/BookStore/db"
	"github.com/XanderMoroz/BookStore/internal/controllers/books"
	"github.com/XanderMoroz/BookStore/internal/controllers/users"
	"github.com/gin-gonic/gin"

	_ "github.com/XanderMoroz/BookStore/docs/bookstore"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

func main() {
	env := config.NewEnv()

	r := gin.Default()

	// Инициализируем базу данных на основе GORM;
	h := db.Init()

	// db.Init()

	// // Регистрируем маршруты приложений
	books.RegisterRoutes(r, h)
	users.RegisterRoutes(r, h)

	// Routes
	r.GET("/", HealthCheck)
	// The url pointing to API definition
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Запускаем сервер на указанном порту
	if err := r.Run(env.AppPort); err != nil {
		log.Fatal(err)
	}

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
