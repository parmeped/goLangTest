module github.com/ginGonicApi/api

go 1.13

require (
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/gin-gonic/contrib v0.0.0-20191209060500-d6e26eeaa607
	github.com/gin-gonic/gin v1.5.0
	github.com/ginGonicApi/repository v0.0.0-00010101000000-000000000000
	github.com/gorilla/sessions v1.2.0 // indirect
)

replace github.com/ginGonicApi/repository => ../repository
