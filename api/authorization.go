package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	repo "github.com/ginGonicApi/repository"
)

const (
	userkey = "testKey"
)

type userLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey) // TODO: [Improvement] this would be the identifier.
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not Authorized!"})
		return
	}

	c.Next()
}

func LoginHandler(data repo.IAuthorizedActions) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// this would come from a Form! Here we parse it from the body
		requestBody := userLoginRequest{}
		c.Bind(&requestBody)

		user := repo.User{
			Name:     requestBody.Name,
			Password: requestBody.Password,
		}

		if strings.Trim(user.Name, " ") == "" || strings.Trim(user.Password, " ") == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
			return
		}

		if !data.ValidateAuthorizedUser(user) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			return
		}

		session.Set(userkey, user.Name) // TODO: [Improvement] first parameter should be the key!
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		message := "Login ok for user " + user.Name
		c.JSON(http.StatusOK, gin.H{"message": message})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
			return
		}
		session.Delete(userkey)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Logout ok!"})
	}
}

func GetCurrentUser(c *gin.Context) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Couldn't get the user from the session!:", err)
		}
	}()

	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return ""
	}
	returnedUser := user.(string)
	return returnedUser
}
