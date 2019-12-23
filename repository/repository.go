package repository

import "github.com/gin-gonic/gin"

type IAuthorizedActions interface {
	PostAuthUser(u User)
	GetAuthUsers() gin.Accounts
	SendMessage(to *User, message Message) bool
	GetUnseenMessages(u *User) []Message
	GetSeenMessages(u *User) []Message
	ValidateAuthorizedUser(u User) bool
	FindUserByName(n string) (*User, string)
	MarkMessagesAsRead(u *User)
}

const (
	SecretUrl = "https://ar.linkedin.com/in/pedro-parmeggiani-a49b6967"
)

type DB struct {
	UsersCollection []User
}

type User struct {
	Name           string
	Password       string
	UnseenMessages []Message
	Messages       []Message
}

type Message struct {
	From string
	To   string
	Body string
}

func New() *DB {
	return initializer()
}

func initializer() *DB {
	var repo = GetRepo()

	var collUsers = []User{}
	collUsers = append(collUsers, repo.Authorized...)

	var db = DB{}
	db.UsersCollection = collUsers

	return &db
}

func GetAuthorized(d *DB) gin.Accounts {
	users := make(map[string]string)
	for _, v := range d.UsersCollection {
		users[v.Name] = v.Password
	}
	var auths gin.Accounts = users
	return auths
}

func (d *DB) FindUserByName(name string) (*User, string) {
	for _, v := range d.UsersCollection {
		if v.Name == name {
			return &v, ""
		}
	}
	err := "Couldn't find the user!"
	return nil, err
}

func (d *DB) MarkMessagesAsRead(u *User) {
	k := 0
	for _, v := range d.UsersCollection {
		if v.Name == u.Name {
			unread := d.UsersCollection[k].UnseenMessages
			d.UsersCollection[k].Messages = append(d.UsersCollection[k].Messages, unread...)
			d.UsersCollection[k].UnseenMessages = nil
			break
		}
		k++
	}
}
