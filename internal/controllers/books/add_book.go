package books

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// Определяем структуру тела запроса на создание экземпляра книги
type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// @Summary        create new Book
// @Description    Creating Book in DB with given request body
// @Tags           Books
// @Accept         json
// @Produce        json
// @Param          request         	body        AddBookRequestBody    true    "Введите данные статьи"
// @Success        201              {string}    string
// @Failure        400              {string}    string    "Bad Request"
// @securityDefinitions.apiKey token
// @in				header
// @name			Authorization
// @Security		JWT
// @Router         /books 			[post]
func (h handler) AddBook(c *gin.Context) {

	body := AddBookRequestBody{}

	// Пытаемся получить тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Создаем новый экземпляр книги
	var book models.Book

	// Присваиваем ему значения из тела запроса и уникальный ID
	book.ID = uuid.New()
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	log.Printf("Добавляем в БД новую книгу:")
	log.Printf("	ID: <%+v>\n", book.ID)
	log.Printf("	Title: <%+v>\n", book.Title)
	log.Printf("	Author: <%+v>\n", book.Author)
	log.Printf("	Description: <%s>\n", book.Description)

	// Пытаемся создать экземпляр книги
	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	c.JSON(http.StatusCreated, &book)
}
