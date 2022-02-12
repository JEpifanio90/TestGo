package main

import (
	"github.com/JEpifanio90/JestGO/controllers"
	"github.com/JEpifanio90/JestGO/models"
	"github.com/gin-gonic/gin"
)

func init() {
	models.ConnectDatabase()
}

func main() {
	engine := gin.Default()

	engine.GET("/books", controllers.FindBooks)
	engine.POST("/books", controllers.CreateBook)
	engine.GET("/books/:id", controllers.FindBook)
	engine.PATCH("/books/:id", controllers.UpdateBook)
	engine.DELETE("/books/:id", controllers.DeleteBook)

	err := engine.Run()

	if err != nil {
		panic(err)
	}
}
