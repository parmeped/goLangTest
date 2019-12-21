package repository

import (
	"github.com/gin-gonic/gin"
)

// TODO: this isn't working. It works when d.MapCollections[0].Data is called.
func (d *DB) PostAuthUser(u User) {
	GetMapCollection(d, "Authorized").Data = append(GetMapCollection(d, "Authorized").Data, u)
}

func (d *DB) GetAuthUsers() gin.Accounts {
	var auths = GetAuthorized(d)
	return auths
}

func (d *DB) SendMessage(message Message) {

}

func (d *DB) GetSeenMessages(u User) *[]Message {
	return &[]Message{}
}

func (d *DB) GetUnseenMessages(u User) *[]Message {
	return &[]Message{}
}
