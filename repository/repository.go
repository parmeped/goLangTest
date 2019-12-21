package repository

import "github.com/gin-gonic/gin"

type IAuthorizedActions interface {
	PostAuthUser(u User)
	GetAuthUsers() gin.Accounts
	SendMessage(message Message)
	GetUnseenMessages(u User) *[]Message
	GetSeenMessages(u User) *[]Message
}

type DB struct {
	Collections    []Collection
	MapCollections []MapCollection
}

type User struct {
	Name           string
	Pass           string
	UnreadMessages []Message
	Messages       []Message
}

type Message struct {
	From    string
	To      string
	Message string
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
	users := make(map[string]string)
	for _, v := range GetMapCollection(d, "Authorized").Data {
		users[v.Name] = v.Pass
	}
	var auths gin.Accounts = users
	return auths
}

// TODO: these are the same! see how to improve, generics would be great
func GetCollection(d *DB, collName string) *Collection {
	for _, v := range d.Collections {
		if v.Name == collName {
			return &v
		}
	}
	return nil
}

func GetMapCollection(d *DB, collName string) *MapCollection {
	for _, v := range d.MapCollections {
		if v.Name == collName {
			return &v
		}
	}
	return nil
}
