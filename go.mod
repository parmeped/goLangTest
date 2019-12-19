module GinGonicApi

go 1.13

require (
	github.com/gin-contrib/sse v0.1.1-0.20190905051334-43f0f29dbd2b
	github.com/gin-gonic/gin v1.5.0
	github.com/ginGonicApi/api v0.0.0-00010101000000-000000000000
	github.com/ginGonicApi/repository v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.3-0.20191022195553-ed6926b37a63
	github.com/json-iterator/go v1.1.9-0.20191112144728-44a7e7340d23
	github.com/mattn/go-isatty v0.0.11
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/ugorji/go v1.1.8-0.20190812104308-42bc974514ff
	golang.org/x/sys v0.0.0-20191218084908-4a24b4065292
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.7
)

replace github.com/ginGonicApi/api => ./api

replace github.com/ginGonicApi/repository => ./repository
