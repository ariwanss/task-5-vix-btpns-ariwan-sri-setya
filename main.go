package main

import (
	"log"
	"net/http"

	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/controllers"
	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/database"
	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/middlewares"
	"github.com/ariwanss/task-5-vix-btpns-ariwan-sri-setya/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}

func loadDb() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Photo{})
}

func main() {
	loadEnv()
	loadDb()

	router := gin.Default()
	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)
	router.PUT("/users/:id", middlewares.Protect, controllers.UpdateUser)
	router.DELETE("/users/:id", middlewares.Protect, controllers.DeleteUser)

	router.DELETE("/users/deleteAll", func(c *gin.Context) {
		res := database.Database.Where("1 = 1").Delete(&models.User{})
		c.JSON(http.StatusOK, gin.H{"rows affected": res.RowsAffected})
	})
	router.GET("/users/all", func(c *gin.Context) {
		var users []models.User
		database.Database.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	router.Run("localhost:5000")
}
