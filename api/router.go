package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repo "github.com/ginGonicApi/repository"
)

func SetupRouter(db *repo.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", LoginHandler())
		v1.GET("/messages", ReadHandler(db))
	}

	authorized := router.Group("/auth", gin.BasicAuth(repo.GetAuthorized(db)))
	{
		authorized.POST("/newAuthorizedUser", PostAuthUserHandler(db))
		authorized.GET("/AuthorizedUsers", GetAuthUsersHandler(db))
		authorized.POST("/sendMessage", SendMessageHandler(db))
	}

	return router
}

func endpointHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "handlerPlaceHolder")
	}
}
