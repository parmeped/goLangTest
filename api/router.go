package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginGonicApi/repository"
)

func SetupRouter(db *repository.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", LoginHandler())
		v1.GET("/read", ReadHandler(db))
	}

	authorized := router.Group("/auth", gin.BasicAuth(repository.GetAuthorised()))
	{
		authorized.POST("/addAuthorizedUser", endpointHandler())
		authorized.POST("/readAuthorizedUsers", endpointHandler())
	}

	return router
}

func endpointHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "handlerPlaceHolder")
	}
}
