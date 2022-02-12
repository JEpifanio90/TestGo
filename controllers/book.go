package controllers

import (
	"github.com/JEpifanio90/JestGO/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(ctx *gin.Context) {
	var books []models.Book

	models.Database.Find(&books)

	ctx.JSON(http.StatusOK, books)
}

func CreateBook(ctx *gin.Context) {
	var ibook models.IBook

	if err := ctx.ShouldBindJSON(&ibook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: ibook.Title, Author: ibook.Author}
	models.Database.Create(&book)

	ctx.JSON(http.StatusCreated, book)
}

func FindBook(ctx *gin.Context) {
	var book models.Book

	if err := models.Database.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func UpdateBook(ctx *gin.Context) {
	var book models.Book

	if err := models.Database.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var ibook models.IBook

	if err := ctx.ShouldBindJSON(&ibook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Database.Model(&book).Updates(ibook)

	ctx.JSON(http.StatusOK, ibook)
}

func DeleteBook(ctx *gin.Context) {
	var book models.Book

	if err := models.Database.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.Database.Delete(&book)

	ctx.JSON(http.StatusNoContent, gin.H{})
}
