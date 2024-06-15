package genres

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// Формируем структуру обработчика базы данных
type handler struct {
	DB *gorm.DB
}

// @Summary        create new genre
// @Description    Creating Genre in DB with given request body
// @Tags           Genges
// @Accept         json
// @Produce        json
// @Param          request         		body        models.CreateGenreBody    true    "Введите название жанра"
// @Success        201              	{string}    string
// @Failure        400              	{string}    string    "Bad Request"
// @Router         /categories 			[post]
func (h handler) CreateNewGenre(c *gin.Context) {
	// body := new(models.CreateGenreBody)
	body := models.CreateGenreBody{}
	// Пытаемся получить тело запроса
	// if err := c.BindJSON(body); err != nil {
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Создаем новый экземпляр книги
	var newGenre models.Genre

	// Присваиваем ему значения из тела запроса и уникальный ID
	newGenre.ID = uuid.New()
	newGenre.Title = body.Title

	log.Printf("Добавляем в БД новый жанр книги:")
	log.Printf("	ID: <%+v>\n", newGenre.ID)
	log.Printf("	Title: <%+v>\n", newGenre.Title)

	// Пытаемся создать экземпляр книги
	if result := h.DB.Create(&newGenre); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	c.JSON(http.StatusCreated, &newGenre)
}

// // @Summary		get all genres
// // @Description Get all genres from db
// // @Tags 		Categories
// // @ID			get-all-genres
// // @Produce		json
// // @Success		200		{object}	[]models.GenreResponse
// // @Router		/categories [get]
// func GetAllGenres(c *fiber.Ctx) error {

// 	categories := utils.GetCategoriesFromDB()

// 	if len(categories) == 0 {
// 		return c.Status(http.StatusNoContent).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "categories not found",
// 			"data":    nil,
// 		})
// 	}

// 	// c.Status(http.StatusOK)
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":  "success",
// 		"message": "Articles Found",
// 		"data":    categories,
// 	})
// }

// // @Summary        add book to category
// // @Description    Adding Article to Category in DB with given request body
// // @Tags           Genre
// // @Accept         json
// // @Produce        json
// // @Param          request         			body        models.AddArticleToCategoryBody    true    "Введите ID статьи и название категории"
// // @Success        201              		{string}    string
// // @Failure        400              		{string}    string    "Bad Request"
// // @Router         /categories/add_article 	[post]
// func AddArticleToCategory(c *fiber.Ctx) error {

// 	db := database.DB

// 	body := new(models.AddArticleToCategoryBody)

// 	// Извлекаем тело запроса
// 	err := c.BodyParser(body)
// 	if err != nil {
// 		// Обрабатываем ошибку
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "Проверьте данные",
// 			"data":    err,
// 		})
// 	}
// 	log.Println("Запрос успешно обработан обработан")

// 	article := utils.GetArticleByIDFromDB(body.ArticleID)

// 	category := utils.GetCategoryByNameFromDB(body.CategoryName)

// 	db.Model(&category).Association("Articles").Append(&article)
// 	db.Model(&article).Association("Categories").Append(&category)

// 	// Возвращаем категорию
// 	return c.Status(201).JSON(fiber.Map{
// 		"status":  "success",
// 		"message": "Article added to Category",
// 		// "data":    newCategory,
// 	})
// }

// // @Summary        delete article from category
// // @Description    Deleting Article from Category in DB with given request body
// // @Tags           Categories
// // @Accept         json
// // @Produce        json
// // @Param          request         			body        models.AddArticleToCategoryBody    true    "Введите ID статьи и название категории"
// // @Success        201              		{string}    string
// // @Failure        400              		{string}    string    "Bad Request"
// // @Router         /categories/remove_article 	[post]
// func DeleteArticleFromCategory(c *fiber.Ctx) error {

// 	db := database.DB

// 	body := new(models.AddArticleToCategoryBody)

// 	// Извлекаем тело запроса
// 	err := c.BodyParser(body)
// 	if err != nil {
// 		// Обрабатываем ошибку
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "Проверьте данные",
// 			"data":    err,
// 		})
// 	}
// 	log.Println("Запрос успешно обработан обработан")

// 	article := utils.GetArticleByIDFromDB(body.ArticleID)

// 	category := utils.GetCategoryByNameFromDB(body.CategoryName)

// 	// result := db.Create(&newCategory)
// 	db.Model(&category).Association("Articles").Delete(&article)
// 	db.Model(&article).Association("Categories").Delete(&category)

// 	// Возвращаем категорию
// 	return c.Status(201).JSON(fiber.Map{
// 		"status":  "success",
// 		"message": "Article delete from Category",
// 	})
// }
