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

type MessagePostRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

// TODO: Check if name or pass are empty. Is this checked somewhere else?
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

// TODO: messages are not being added to the unseen ones
func SendMessageHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := MessagePostRequest{}
		c.Bind(&requestBody)

		message := repo.Message{
			From: requestBody.From,
			To:   requestBody.To,
			Body: requestBody.Body,
		}

		if message.From == message.To {
			c.JSON(http.StatusBadRequest, "Can't send a message to oneself!")
			return
		}

		if u, err := data.FindUserByName(message.To); err != "" {
			c.JSON(http.StatusInternalServerError, err)
			return
		} else {
			if data.SendMessage(u, message) {
				c.JSON(http.StatusNoContent, "Message sent successfully")
			} else {
				c.JSON(http.StatusInternalServerError, "There was an error sending the message!")
			}
		}
	}
}

func UnseenMessagesHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := GetCurrentUser(c)

		if u, err := data.FindUserByName(userName); err != "" {
			c.JSON(http.StatusInternalServerError, err)
			return
		} else {
			if uMessages := data.GetUnseenMessages(u); len(uMessages) > 0 {
				c.JSON(http.StatusOK, uMessages)
				data.MarkMessagesAsRead(u)
				return
			} else {
				c.JSON(http.StatusOK, "User has no unseen messages")
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
			c.JSON(http.StatusInternalServerError, err)
			return
		} else {
			if messages := data.GetSeenMessages(u); len(messages) > 0 {
				c.JSON(http.StatusOK, messages)
				return
			} else {
				c.JSON(http.StatusOK, "User has no Seen messages")
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
