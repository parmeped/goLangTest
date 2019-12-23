package main

import (
	"goApi/http/handler"
	"goApi/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/getFeed", handler.NewsFeedGet(feed))
	r.POST("/postFeed", handler.NewsFeedPost(feed))

	r.Run()
}
