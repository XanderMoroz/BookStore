package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/config"
	"github.com/XanderMoroz/BookStore/db"

	"github.com/XanderMoroz/BookStore/internal/controllers/books"
	"github.com/XanderMoroz/BookStore/internal/controllers/users"

	_ "github.com/XanderMoroz/BookStore/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			BookStore APIs
//	@version		1.0
//	@description	Testing Swagger APIs.
//	@termsOfService	http://swagger.io/terms/
// 	@contact.name API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
// 	@securityDefinitions.apiKey JWT
//	@in				header
//	@name			token//
//	@license.name Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html//
//	@host localhost:8082
//	@BasePath		/api/v1/

func main() {
	env := config.NewEnv()

	r := gin.Default()

	// Инициализируем базу данных на основе GORM;
	h := db.Init()

	// Routes
	r.GET("/", HealthCheck)

	// The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// // Регистрируем маршруты приложений
	books.RegisterRoutes(r, h)
	users.RegisterRoutes(r, h)
	// Запускаем сервер на указанном порту
	if err := r.Run(":" + env.AppPort); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
//
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			root
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"success":  "True",
		"message":  "Server is up and running",
		"docs_url": "http://127.0.0.1:8082/swagger/index.html",
		"error":    "False",
	}

	c.JSON(http.StatusOK, res)
}
