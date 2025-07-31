package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loid-lab/restaurant-menu/initializers"
	"github.com/loid-lab/restaurant-menu/models"
)

func GetAllInvoices(c *gin.Context) {
	var invoices []models.Invoice
	if err := initializers.DB.Preload("item").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"invoices": invoices})
}
