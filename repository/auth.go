package repository

import "github.com/gin-gonic/gin"

func AddAuthorizedUser(u gin.Accounts) bool {
	return true
}

func ReadAuthorizedUsers() gin.Accounts {
	return GetAuthorised()
}
