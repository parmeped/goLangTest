package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := setupRouter()

	r.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint())
		v1.POST("/submit", submitEndpoint())
		v1.POST("/read", readEndpoint())
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint2())
		v2.POST("/submit", submitEndpoint2())
		v2.POST("/read", readEndpoint2())
	}

	return router
}

func loginEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "loginEndpointV1")
	}
}

func submitEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "submitEndpointV1")
	}
}

func readEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "readEndpointV1")
	}
}

func loginEndpoint2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "loginEndpointV2")
	}
}

func submitEndpoint2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "submitEndpointV2")
	}
}

func readEndpoint2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "readEndpointV2")
	}
}
