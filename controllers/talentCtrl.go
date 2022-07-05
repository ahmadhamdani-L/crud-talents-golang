package controllers

import (
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TalentInput struct {
	Id         string `json:"id"`
	NamaTalent string `json:"nama_talent"`
}

//tampil data talents
func TalentTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var talent []models.Talent
	db.Find(&talent)
	c.JSON(http.StatusOK, gin.H{"data": talent})
}

func TalentAdd(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dataInput TalentInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//proses input
	talent := models.Talent{
		Id:         dataInput.Id,
		NamaTalent: dataInput.NamaTalent,
	}
	db.Create(&talent)
	c.JSON(http.StatusOK, gin.H{"data": talent})
}
