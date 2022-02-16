package controllers

import (
	"fmt"
	"github.com/JEpifanio90/JestGO/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "jester"
const title = "Auth"

func Register(ctx *gin.Context) {
	var rawUser models.IUser

	if err := ctx.ShouldBindJSON(&rawUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"title": fmt.Sprintf("%s-Register", title), "message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(rawUser.Password), 12)

	user := models.User{
		Name:     rawUser.Name,
		Email:    rawUser.Email,
		Password: password,
	}

	models.Database.Create(&user)

	ctx.JSON(http.StatusCreated, user)
}

func Login(ctx *gin.Context) {
	var credentials models.ICredentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"title": fmt.Sprintf("%s-Login", title), "message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	var user models.User

	models.Database.Where("email = ?", credentials.Email).First(&user)

	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"title": fmt.Sprintf("%s-Login", title), "message": "Invalid Credentials", "status": http.StatusNotFound})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(credentials.Password)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"title": fmt.Sprintf("%s-Login", title), "message": "Invalid Credentials", "status": http.StatusNotFound})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{Issuer: strconv.Itoa(int(user.Id)), ExpiresAt: time.Now().Add(time.Hour * 1).Unix()})
	rClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{Issuer: strconv.Itoa(int(user.Id)), ExpiresAt: time.Now().Add(time.Hour * 12).Unix()})

	token, err := claims.SignedString([]byte(SecretKey))
	rToken, rErr := rClaims.SignedString([]byte(SecretKey))

	if err != nil || rErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"title": fmt.Sprintf("%s-Login", title), "message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": rToken, "expires_at": time.Now().Add(time.Hour * 1).UnixMilli()})
}
