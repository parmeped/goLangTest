package repository

import "github.com/gin-gonic/gin"

var authorised = gin.Accounts{
	"TestUser": "pass",
}

var testArray = [3]string{
	"test1",
	"test2",
	"test3",
}

type Repo struct {
	Authorised map[string]string
	TestArray  [3]string
}

func GetRepo() *Repo {
	repo := Repo{authorised, testArray}
	return &repo
}

func GetAuthorised() gin.Accounts {
	return authorised
}
