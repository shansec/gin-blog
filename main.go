package main

import (
	"myginblog/model"
	"myginblog/routes"
)

func main() {

	model.InitDb()
	routes.InitRoute()
}
