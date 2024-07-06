package handlers

import (
	"crud/go/database"
	"crud/go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := database.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}
