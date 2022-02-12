package controllers

import (
	"github.com/JEpifanio90/JestGO/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(context *gin.Context) {
	var books []models.Book

	models.Database.Find(&books)

	context.JSON(http.StatusOK, gin.H{"data": books})
}
