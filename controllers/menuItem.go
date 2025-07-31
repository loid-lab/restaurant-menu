package controllers

import (
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/loid-lab/restaurant-menu/initializers"
	"github.com/loid-lab/restaurant-menu/models"
)

func CreateMenuItem(c *gin.Context) {
	var menuItem models.MenuItem

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too big"})
		return
	}

	file, _, err := c.Request.FormFile("image")
	if err == nil && file != nil {
		uploadParams := uploader.UploadParams{Folder: "menu_items"}
		uploadResult, err := initializers.Cloudinary.Upload.Upload(c, file, uploadParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Image upload failed"})
			return
		}
		menuItem.ImageURL = uploadResult.SecureURL
	}

	menuItem.Name = c.PostForm("name")
	categoryID := c.PostForm("cetegory_id")
	if categoryID != "" {
		var id uint
		fmt.Sscan(categoryID, "%d", &id)
		menuItem.CategoryID = id
	}

	user, _ := c.Get("currentUser")
	menuItem.CreatedBy = user.(models.User).ID

	if err := initializers.DB.Create(&menuItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create menu item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menu_item": menuItem})
}

func GetAllMenuItems(c *gin.Context) {
	var menuItems []models.MenuItem

	if err := initializers.DB.Find(&menuItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch menu items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menu_items": menuItems})
}

func UpdateMenuItems(c *gin.Context) {
	id := c.Param("id")
	var allMenuItems models.MenuItem

	if err := c.ShouldBindJSON(&allMenuItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var menuItem models.MenuItem
	initializers.DB.First(&menuItem, id)

	if menuItem.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Menu item not found"})
		return
	}

	initializers.DB.Model(&menuItem).Updates(allMenuItems)
	c.JSON(http.StatusOK, gin.H{
		"menu_item": menuItem})
}

func GetMenuItemsByID(c *gin.Context) {
	id := c.Param("id")
	var menuItem models.MenuItem

	if err := initializers.DB.First(&menuItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Menu item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menu_item": menuItem})
}

func DeleteMenuItem(c *gin.Context) {
	id := c.Param("id")

	var menuItem models.MenuItem
	initializers.DB.First(&menuItem, id)

	if menuItem.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Menu item not found"})
		return
	}

	initializers.DB.Delete(&menuItem)
	c.JSON(http.StatusOK, gin.H{
		"message": "Menu item deleted"})
}
