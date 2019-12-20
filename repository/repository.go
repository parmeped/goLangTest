package repository

import "github.com/gin-gonic/gin"

type IAddAuthorizedUser interface {
	AddAuthorizedUser(u gin.Accounts) bool
}

type IReadAuthorizedUsers interface {
	ReadAuthorizedUsers() gin.Accounts
}

type DB struct {
	Collections [1]Collection
}

type Collection struct {
	Name string
	Data [3]string
}

func New() *DB {
	return initializer()
}

func initializer() *DB {
	var repo = GetRepo()

	var coll = Collection{"testing", repo.TestArray}
	var collections [1]Collection
	collections[0] = coll

	var db = DB{collections}
	return &db
}
