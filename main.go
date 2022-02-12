package main

import (
	"fmt"
	"github.com/JEpifanio90/JestGO/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	fmt.Println("Trying to connect to db")
	models.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
