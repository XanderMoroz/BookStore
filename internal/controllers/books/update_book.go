package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// @Summary			update book by ID
// @Description 	Update book by ID
// @ID				delete-book-by-id
// @Tags 			Books
// @Produce			json
// @Param			id					path		int								true	"Book ID"
// @Param           request         	body        models.UpdateBookRequestBody    true    "Введите новые данные книги"
// @Success			200	{object}	[]string
// @Failure			404	{object}	[]string
// @Router			/books/{id} 	[put]
func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	body := models.UpdateBookRequestBody{}

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
