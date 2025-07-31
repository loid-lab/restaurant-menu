package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loid-lab/restaurant-menu/initializers"
	"github.com/loid-lab/restaurant-menu/models"
)

func AddToCart(c *gin.Context) {
	var input struct {
		MenuItemID uint
		Quantity   int
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user, _ := c.Get("currentUser")

	cartItem := models.CartItem{
		UserID:     user.(models.User).ID,
		MenuItemID: input.MenuItemID,
		Quantity:   input.Quantity,
	}

	if err := initializers.DB.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not add to carr"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_item": cartItem})
}

func GetCart(c *gin.Context) {
	user, _ := c.Get("currentUser")

	var cartItem []models.CartItem
	initializers.DB.Where("user_id=?", user.(models.User).ID).First(&cartItem)
	c.JSON(http.StatusOK, gin.H{
		"cart": cartItem})
}

func DeleteCartItem(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("currentUser")

	var cartItem models.CartItem
	if err := initializers.DB.First(&cartItem, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cart item not found"})
		return
	}

	if cartItem.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access to cart item"})
		return
	}

	initializers.DB.Delete(&cartItem)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cart item deleted"})
}
