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
			Name:     requestBody.Name,
			Password: requestBody.Password,
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

// TODO: finish this method!
func SendMessageHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "authHandlerPlaceHolder")
	}
}

func UnseenMessagesHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if logged in, we can assume the user exists!
		userName := GetCurrentUser(c)

		if u, err := data.FindUserByName(userName); err != "" {
			c.JSON(http.StatusInternalServerError, "Couldn't retrieve the user!")
			return
		} else {
			if u.UnseenMessages != nil {
				c.JSON(http.StatusOK, u.UnseenMessages)
				data.MarkMessagesAsRead(u)
				return
			} else {
				c.JSON(http.StatusOK, "User has no unseen Messages")
				return
			}
		}
	}
}

func SeenMessagesHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if logged in, we can assume the user exists!
		userName := GetCurrentUser(c)

		if u, err := data.FindUserByName(userName); err != "" {
			c.JSON(http.StatusInternalServerError, "Couldn't retrieve the user!")
			return
		} else {
			if u.Messages != nil {
				c.JSON(http.StatusOK, u.Messages)
				return
			} else {
				c.JSON(http.StatusOK, "User has no Messages")
				return
			}
		}
	}
}

func SecretRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, repo.SecretUrl)
		return
	}
}
