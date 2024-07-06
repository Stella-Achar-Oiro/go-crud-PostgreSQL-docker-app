package main

import (
	"crud/go/database"
	"crud/go/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()

	r.POST("/items", handlers.CreateItem)
	r.GET("/items/:id", handlers.GetItem)
	r.PUT("/items/:id", handlers.UpdateItem)
	r.DELETE("/items/:id", handlers.DeleteItem)

	r.Run(":8080")
}
