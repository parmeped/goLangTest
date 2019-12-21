package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repo "github.com/ginGonicApi/repository"
)

func SetupRouter(db *repo.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/")
	{
		v1.POST("/login", LoginHandler())
	}

	authorized := router.Group("/auth", gin.BasicAuth(repo.GetAuthorized(db)))
	{
		authorized.POST("/newAuthorizedUser", PostAuthUserHandler(db))
		authorized.GET("/AuthorizedUsers", GetAuthUsersHandler(db))
		authorized.POST("/sendMessage", SendMessageHandler(db))
		authorized.GET("/seenMessages", SeenMessagesHandler(db))
		authorized.GET("/unseenMessages", UnseenMessagesHandler(db))
	}

	return router
}

func endpointHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "handlerPlaceHolder")
	}
}
