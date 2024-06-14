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

// @Summary		get all my orders
// @Description Get all my orders from db
// @Tags 		Orders
// @ID			get-all-my-orders
// @Produce		json
// @Success		200				{object}	[]models.OrderResponse
// @Router		/user/orders 	[get]
func (h handler) GetMyOrders(c *gin.Context) {

	user_id, err := utils.ExtractUserIDFromToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := utils.GetUserByIDFromDB(user_id)

	// Создаем пустой срез экземпляров
	var orders []models.Order

	// Пытаемся найти экземпляры
	if result := h.DB.Preload("User").Preload("Items").Find(&orders, "user_id = ?", currentUser.ID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	} else {
		log.Println("Удалось извлеч")
	}
	// Отправляем в контекст список экземпляров
	c.JSON(http.StatusOK, &orders)
}

// @Summary		get my order by ID
// @Description Get my order by ID
// @Tags 		Orders
// @ID			get-order-by-id
// @Produce		json
// @Param		id				path		string					true	"Order ID"
// @Success		200				{object}	models.OrderResponse
// @Failure		404				{object}	[]string
// @Router		/user/orders/{id} 	[get]
func (h handler) GetMyOrder(c *gin.Context) {

	user_id, err := utils.ExtractUserIDFromToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := utils.GetUserByIDFromDB(user_id)

	// Извлекаем из контекста значение параметра ID
	id := c.Param("id")

	// Инициализируем пустой экземпляр книги
	var order models.Order

	// Пытаемся найти экземпляр заказа с полученным значением ID
	if result := h.DB.Preload("User").Preload("Items").First(&order, "user_id = ? AND id >= ?", currentUser.ID, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Order Not Found",
		})
		return
	} else {
		log.Println("Заказ - успешно извлечен:")
		log.Printf("	Заказ ID: <%v>\n", order.ID)
		log.Printf("	Имя Заказчика: <%v>\n", order.User.Name)
	}

	// Отправляем в контекст экземпляр заказа
	c.JSON(http.StatusOK, &order)
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

	user_id, err := utils.ExtractUserIDFromToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := utils.GetUserByIDFromDB(user_id)

	// Извлекаем из контекста значение параметра ID
	id := c.Param("id")

	// Инициализируем пустой экземпляр книги
	var order models.Order

	// Пытаемся найти экземпляры
	if result := h.DB.First(&order, "user_id = ? AND id >= ?", currentUser.ID, id); result.Error != nil {
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
