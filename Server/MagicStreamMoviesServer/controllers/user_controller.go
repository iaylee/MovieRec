package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iaylee/MovieRec/Server/MagicStreamMoviesServer/models"
)

func RegisterUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var user models.User
		if err:=c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input data"})
			return
		}
		validate := validator.New()
		if err := validate.Struct(user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":"Validation failed", "details":err.Error()})
			return
		}

		
	}
}