package handlers

import (
	"crud/go/database"
	"crud/go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := database.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}
