package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/helpers"
	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/models"
	"github.com/gin-gonic/gin"
)

var (
	InvalidPhotoId = errors.New("invalid photo id")
)

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	err := c.ShouldBindJSON(&photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authPayload := c.Value("User").(*helpers.AuthPayload)
	photo.SetUserId(authPayload.UserId)

	savedPhoto, err := models.CreatePhoto(&photo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedPhoto)
}

func GetPhoto(c *gin.Context) {
	user := c.Value("User").(*helpers.AuthPayload)
	photo, err := models.FindPhotoByUserId(user.UserId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func UpdatePhoto(c *gin.Context) {
	idStr, ok := c.Params.Get("id")

	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Photo id not provided"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid photo id"})
		return
	}

	photo, err := models.FindPhotoById(uint(id))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	authPayload := c.Value("User").(*helpers.AuthPayload)

	if authPayload.UserId != photo.UserId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized wrong user"})
		return
	}

	var photoUpdate models.Photo

	if err = c.ShouldBindJSON(&photoUpdate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPhoto, err := models.UpdatePhoto(uint(id), &photoUpdate)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPhoto)
}

func DeletePhoto(c *gin.Context) {
	idStr, ok := c.Params.Get("id")
	
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Photo id not provided"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid photo id"})
		return
	}

	photo, err := models.FindPhotoById(uint(id))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	authPayload := c.Value("User").(*helpers.AuthPayload)

	if authPayload.UserId != photo.UserId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized wrong user"})
		return
	}

	err = models.DeletePhoto(uint(id))
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
