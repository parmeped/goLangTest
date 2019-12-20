package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repo "github.com/ginGonicApi/repository"
)

type AuthUserPostRequest struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

func PostAuthUserHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := AuthUserPostRequest{}
		c.Bind(&requestBody)

		user := repo.User{
			Name: requestBody.Name,
			Pass: requestBody.Password,
		}

		data.PostAuthUser(user)

		c.Status(http.StatusNoContent)

	}
}

func GetAuthUsersHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var auths = data.GetAuthUsers()
		c.JSON(http.StatusOK, auths)
	}
}

func SendMessageHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "authHandlerPlaceHolder")
	}
}
