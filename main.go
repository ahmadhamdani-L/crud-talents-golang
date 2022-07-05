package main

import (
	"golang-api/controllers"
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//models
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "danslelouch"})
	})

	r.GET("/talent", controllers.TalentTampil)
	r.POST("/talent/add", controllers.TalentAdd)
	r.Run()

}
