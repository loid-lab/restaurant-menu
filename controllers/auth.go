package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/loid-lab/restaurant-menu/initializers"
	"github.com/loid-lab/restaurant-menu/models"
	"github.com/loid-lab/restaurant-menu/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email=?", userInput.Email).First(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already in use"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user := models.User{
		Email:    userInput.Email,
		Password: string(passwordHash),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User"})
		return
	}

	utils.SendMail(models.EmailData{
		To:       user.Email,
		Subject:  "Welcome to (insert restaurant name here)",
		HTMLBody: fmt.Sprintf("<h2>Welcome %s!</h2><p>Your account has been created</p>", user.Email),
	})
}

func Login(c *gin.Context) {
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("email=?", userInput.Email).First(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInput.Password), []byte(userFound.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr})
}
