package repository

import "github.com/gin-gonic/gin"

type IAuthorizedActions interface {
	PostAuthUser(u User)
	GetAuthUsers() gin.Accounts
	SendMessage() bool
}

type DB struct {
	Collections    []Collection
	MapCollections []MapCollection
}

type User struct {
	Name string
	Pass string
}

type Collection struct {
	Name string
	Data []string
}

type MapCollection struct {
	Name string
	Data []User
}

func New() *DB {
	return initializer()
}

func initializer() *DB {
	var repo = GetRepo()

	var coll = Collection{Name: "Testing"}

	coll.Data = append(coll.Data, repo.TestArray...)

	var collAuth = MapCollection{Name: "Authorized"}
	collAuth.Data = append(collAuth.Data, repo.Authorized...)

	var db = DB{}
	db.Collections = append(db.Collections, coll)
	db.MapCollections = append(db.MapCollections, collAuth)

	return &db
}

func GetAuthorized(d *DB) gin.Accounts {
	// TODO we are assuming MapCollections[0] is authorized users!!
	users := make(map[string]string)
	for _, v := range d.MapCollections[0].Data {
		users[v.Name] = v.Pass
	}
	var auths gin.Accounts = users
	return auths
}
