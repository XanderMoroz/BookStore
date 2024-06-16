package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// @Summary		get all books
// @Description Get all books from db
// @Tags 		Books
// @ID			get-all-books
// @Produce		json
// @Success		200		{object}	[]models.BookResponse
// @Router		/books [get]
func (h handler) GetBooks(c *gin.Context) {

	// Создаем пустой срез экземпляров
	var books []models.Book

	// Пытаемся найти экземпляры
	if result := h.DB.Preload("Genres").Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Отправляем в контекст список экземпляров
	c.JSON(http.StatusOK, &books)
}
