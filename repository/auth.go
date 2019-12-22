package repository

import (
	"github.com/gin-gonic/gin"
)

// --------------- Interface implementation ---------------

func (d *DB) PostAuthUser(u User) {
	d.UsersCollection = append(d.UsersCollection, u)
}

// this awesome method reveals all users passwords! :D
func (d *DB) GetAuthUsers() gin.Accounts {
	var auths = GetAuthorized(d)
	return auths
}

func (d *DB) SendMessage(to *User, message Message) bool {
	for _, v := range d.UsersCollection {
		if v.Name == to.Name {
			v.UnseenMessages = append(v.UnseenMessages, message)
			return true
		}
	}
	return false
}

func (d *DB) GetSeenMessages(u *User) []Message {
	messages := []Message{}
	for _, v := range d.UsersCollection {
		if v.Name == u.Name {
			messages = v.Messages
		}
	}
	return messages
}

func (d *DB) GetUnseenMessages(u *User) []Message {
	uMessages := []Message{}
	for _, v := range d.UsersCollection {
		if v.Name == u.Name {
			uMessages = v.UnseenMessages
		}
	}
	return uMessages
}

func (d *DB) ValidateAuthorizedUser(u User) bool {
	for _, v := range d.UsersCollection {
		if v.Name == u.Name && v.Password == u.Password {
			return true
		}
	}
	return false
}
