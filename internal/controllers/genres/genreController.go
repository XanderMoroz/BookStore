package genres

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/XanderMoroz/BookStore/db"
	"github.com/XanderMoroz/BookStore/internal/models"
	"github.com/XanderMoroz/BookStore/internal/utils"
)

// @Summary        create new genre
// @Description    Creating Genre in DB with given request body
// @Tags           Genres
// @Accept         json
// @Produce        json
// @Param          request         		body        models.CreateGenreBody    true    "Введите название жанра"
// @Success        201              	{string}    string
// @Failure        400              	{string}    string    "Bad Request"
// @Router         /genres 				[post]
func (h handler) CreateNewGenre(c *gin.Context) {

	body := new(models.CreateGenreBody)
	// Пытаемся получить тело запроса
	if err := c.BindJSON(body); err != nil {
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

// @Summary		get all genres
// @Description Get all genres from db
// @Tags 		Genres
// @ID			get-all-genres
// @Produce		json
// @Success		200		{object}	[]models.GenreResponse
// @Router		/genres [get]
func (h handler) GetAllGenres(c *gin.Context) {

	genres := utils.GetAllGenresFromDB()

	if len(genres) == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, &genres)
}

// @Summary        add book to genre
// @Description    Adding Book to Genre in DB with given request body
// @Tags           Genres
// @Accept         json
// @Produce        json
// @Param          request         			body        models.BookToGenreBody    true    "Введите ID книги и название жанра"
// @Success        201              		{string}    string
// @Failure        400              		{string}    string    "Bad Request"
// @Router         /genres/add_book 		[post]
func (h handler) AddBookToGenre(c *gin.Context) {

	db := db.DB

	body := new(models.BookToGenreBody)

	// Извлекаем тело запроса
	if err := c.BindJSON(body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println("Запрос успешно обработан обработан")

	book := utils.GetBookByIDFromDB(body.BookID)

	genre := utils.GetGenreByNameFromDB(body.Genre)

	db.Model(&genre).Association("Books").Append(&book)
	db.Model(&book).Association("Genres").Append(&genre)

	// Возвращаем книгу
	c.JSON(http.StatusOK, &book)
}

// @Summary        delete book from genre
// @Description    Deleting Book from Genre in DB with given request body
// @Tags           Genres
// @Accept         json
// @Produce        json
// @Param          request         			body        models.BookToGenreBody    true    "Введите ID Книги и название жанра"
// @Success        201              		{string}    string
// @Failure        400              		{string}    string    "Bad Request"
// @Router         /genres/remove_book 	[post]
func (h handler) RemoveBookFromGenre(c *gin.Context) {

	db := db.DB

	body := new(models.BookToGenreBody)

	// Извлекаем тело запроса
	if err := c.BindJSON(body); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Bad Request Check your data",
			"data":    err,
		})
		// c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println("Запрос успешно обработан обработан")

	book := utils.GetBookByIDFromDB(body.BookID)

	genre := utils.GetGenreByNameFromDB(body.Genre)

	// result := db.Create(&newCategory)
	db.Model(&genre).Association("Books").Delete(&book)
	db.Model(&book).Association("Genres").Delete(&genre)

	// Возвращаем категорию
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book remove from Genre",
		"data":    &book,
	})
}
