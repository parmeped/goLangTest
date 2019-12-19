package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginGonicApi/repository"
)

func SetupRouter(db *repository.DB) *gin.Engine {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", LoginHandler())
		v1.POST("/submit", SubmitHandler())
		v1.POST("/read", ReadHandler(db))
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", endpointHandler())
		v2.POST("/submit", endpointHandler())
		v2.POST("/read", endpointHandler())
	}

	return router
}

func endpointHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "handlerPlaceHolder")
	}
}
