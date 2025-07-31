package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/loid-lab/restaurant-menu/initializers"
	"github.com/loid-lab/restaurant-menu/models"
	"github.com/loid-lab/restaurant-menu/utils"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user, _ := c.Get("currentUser")
	order.UserID = user.(models.User).ID

	if err := initializers.DB.Create(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create order"})
		return
	}

	invoice := models.Invoice{
		InvoiceNumber: fmt.Sprintf("INV-%d", time.Now().Unix()),
		Date:          time.Now(),
		CustomerName:  user.(models.User).FullName,
		Items:         []models.InvoiceItem{},
	}

	for _, item := range order.Items {
		invoice.Items = append(invoice.Items, models.InvoiceItem{
			MenuItemName: item.MenuItem.Name,
			Quantity:     item.Quantity,
			UnitPrice:    item.UnitPrice,
			TotalPrice:   item.TotalPrice,
		})
	}

	invoice.UserID = order.UserID

	if err := initializers.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create invoice"})
		return
	}

	email := models.EmailData{
		From:     "no-reply@gamil.com",
		To:       user.(models.User).Email,
		Subject:  "Your (restaurant name) Invoice",
		HTMLBody: "<p>Thanks for your order! YOur invoice is attached</p>",
		SMTConfig: models.SMTConfig{
			SMTPHost: initializers.Env.SMTPHost,
			SMTPPort: initializers.Env.SMTPPort,
			SMTPUser: initializers.Env.SMTPUser,
			SMTPPass: initializers.Env.SMTPPass,
		},
	}

	go utils.GenerateSendInvoice(invoice, email)

	c.JSON(http.StatusOK, gin.H{
		"order": order})
}

func GetUserOrder(c *gin.Context) {
	var orders []models.Order

	if err := initializers.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch menu items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders})
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := initializers.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order})
}
