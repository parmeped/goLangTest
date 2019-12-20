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
		c.JSON(http.StatusOK, db.Collections[0].Data)
	}

}
