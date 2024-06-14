package orders

import (
	"log"
	"net/http"

	"github.com/XanderMoroz/BookStore/internal/models"
	"github.com/XanderMoroz/BookStore/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Формируем структуру обработчика базы данных
type handler struct {
	DB *gorm.DB
}

// @Summary        create new order
// @Description    Creating Order in DB with given request body
// @Tags           Orders
// @Accept         json
// @Produce        json
// @Success        201              		{string}    string
// @Failure        400              	{string}    string    "Bad Request"
// @securityDefinitions.apiKey token
// @in				header
// @name			Authorization
// @Security		JWT
// @Router         /user/orders 			[post]
func (h handler) AddOrder(c *gin.Context) {

	user_id, err := utils.ExtractUserIDFromToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := utils.GetUserByIDFromDB(user_id)

	// Создаем новый экземпляр заказа
	var newOrder models.Order

	// Присваиваем ему значения из тела запроса и уникальный ID
	newOrder.ID = uuid.New()
	newOrder.UserID = currentUser.ID
	newOrder.User = currentUser

	log.Printf("Добавляем в БД новый пустой заказ:")
	log.Printf("	ID: <%+v>\n", newOrder.ID)
	log.Printf("	Создан пользователем : <%+v>\n", newOrder.User)

	// Пытаемся создать экземпляр книги
	if result := h.DB.Create(&newOrder); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	c.JSON(http.StatusCreated, &newOrder)
}

// @Summary		get an book by ID
// @Description Get an book by ID
// @Tags 		Books
// @ID			get-book-by-id
// @Produce		json
// @Param		id				path		string					true	"Book ID"
// @Success		200				{object}	models.BookResponse
// @Failure		404				{object}	[]string
// @Router		/books/{id} 	[get]
func (h handler) GetBook(c *gin.Context) {

	// Извлекаем из контекста значение параметра ID
	id := c.Param("id")

	// Инициализируем пустой экземпляр книги
	var book models.Book

	// Пытаемся найти экземпляр книги с полученным значением ID
	if result := h.DB.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book Not Found",
		})
		return
	} else {
		log.Println("Книга - успешно извлечена:")
		log.Printf("Author: <%v>\n", book.Author)
		log.Printf("Title: <%v>\n", book.Title)
		log.Printf("Description: <%v>\n", book.Description)
	}

	// Отправляем в контекст экземпляр книги
	c.JSON(http.StatusOK, &book)
}

// @Summary		delete a order by ID
// @Description Delete a order by ID
// @ID			delete-order-by-id
// @Tags 		Orders
// @Produce		json
// @Param		id					path		string		true	"Order ID"
// @Success		200					{object}	[]string
// @Failure		404					{object}	[]string
// @Router		/user/orders/{id} 	[delete]
func (h handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")
	var order models.Order

	// Пытаемся найти экземпляры
	if result := h.DB.First(&order, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   result.Error,
		})
		return
	}

	h.DB.Delete(&order)

	c.JSON(http.StatusNotFound, gin.H{
		"success": true,
		"message": "Order Deleted",
	})

}
