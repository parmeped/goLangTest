package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	repo "github.com/ginGonicApi/repository"
)

func SetupRouter(db *repo.DB) *gin.Engine {
	router := gin.Default()

	router.Use(sessions.Sessions("UserSessionCookie", sessions.NewCookieStore([]byte("secret"))))

	router.POST("/login", LoginHandler(db))

	authorized := router.Group("/auth")
	authorized.Use(AuthRequired) // here we validate user session
	{
		authorized.GET("/authorizedUsers", GetAuthUsersHandler(db))
		authorized.GET("/seenMessages", SeenMessagesHandler(db))
		authorized.GET("/unseenMessages", UnseenMessagesHandler(db))
		authorized.POST("/newAuthorizedUser", PostAuthUserHandler(db))
		authorized.POST("/sendMessage", SendMessageHandler(db))
		authorized.POST("/logout", LogoutHandler())
	}

	router.GET("/superSecretRoute", SecretRouteHandler())
	return router
}

// func endpointHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "handlerPlaceHolder")
// 	}
// }
