package repository

import (
	"github.com/gin-gonic/gin"
)

func (d *DB) PostAuthUser(u User) {
	d.UsersCollection = append(d.UsersCollection, u)
}

// this awesome method reveals all users passwords! :D
func (d *DB) GetAuthUsers() gin.Accounts {
	var auths = GetAuthorized(d)
	return auths
}

// TODO: implement sendMessage
func (d *DB) SendMessage(message Message) {

}

func (d *DB) GetSeenMessages(u User) *[]Message {
	return &[]Message{}
}

func (d *DB) GetUnseenMessages(u User) *[]Message {
	return &[]Message{}
}

func (d *DB) ValidateAuthorizedUser(u User) bool {
	for _, v := range d.UsersCollection {
		if v.Name == u.Name && v.Password == u.Password {
			return true
		}
	}
	return false
}
