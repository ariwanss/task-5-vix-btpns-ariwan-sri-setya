package controllers

import (
	"net/http"
	"strconv"

	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/helpers"
	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := models.CreateUser(&newUser)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := helpers.GenerateToken(savedUser.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  savedUser,
		"token": token,
	})
}

func Login(c *gin.Context) {
	var attemptingUser models.User

	if err := c.ShouldBindJSON(&attemptingUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(attemptingUser.Username)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = helpers.ComparePassword(attemptingUser.Password, user.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := helpers.GenerateToken(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	idStr, exists := c.Params.Get("id")

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id not provided"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authPayload := c.Value("User").(*helpers.AuthPayload)
	authUserId := authPayload.UserId

	if uint(id) != authUserId {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unauthorized wrong owner"})
		return
	}

	var userUpdate models.User

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := models.UpdateUser(uint(id), &userUpdate)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	idStr, exists := c.Params.Get("id")

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id not provided"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authPayload := c.Value("User").(*helpers.AuthPayload)
	authUserId := authPayload.UserId

	if uint(id) != authUserId {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unauthorized wrong owner"})
		return
	}

	err = models.DeleteUser(uint(id))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
