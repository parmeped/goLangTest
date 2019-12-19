package main

import (
	"github.com/ginGonicApi/api"
	"github.com/ginGonicApi/repository"
)

func main() {
	db := repository.New()
	r := api.SetupRouter(db)

	r.Run()
}
