package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginGonicApi/repository"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "handlerPlaceHolder")
	}
}

func SubmitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "handlerPlaceHolder")
	}
}

func ReadHandler(db *repository.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = [3]string
		for i := 0; i < len(db.Tables[0].Data); i++ {
			data[i] = db.Tables[0].Data[i]
		}
		c.JSON(http.StatusOK, data)
	}
	
}
