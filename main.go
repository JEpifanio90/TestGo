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
	r := gin.Default()

	r.GET("/books", controllers.FindBooks)

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
