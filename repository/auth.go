package repository

import (
	"github.com/gin-gonic/gin"
)

func (d *DB) PostAuthUser(u User) {
	// TODO here we are again assuming MapCollections[0].Data is authorized users
	d.MapCollections[0].Data = append(d.MapCollections[0].Data, u)
}

func (d *DB) GetAuthUsers() gin.Accounts {
	var auths = GetAuthorized(d)
	return auths
}

func (d *DB) SendMessage() bool {
	return true
}
