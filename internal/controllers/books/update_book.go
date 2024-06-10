package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// Определяем структуру тела запроса на создание экземпляра книги
type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// Определяем функцию обработчик для создания нового экземпляра
func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	body := UpdateBookRequestBody{}

	// Пытаемся получить тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Создаем новый экземпляр книги
	var book models.Book

	// Пытаемся найти экземпляр книги с полученным значением ID
	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Присваиваем ему значения из тела запроса
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	// Сохраняем изменения в базе данных
	h.DB.Save(&book)

	// Отправляем в контекст экземпляр книги
	c.JSON(http.StatusOK, &book)
}
