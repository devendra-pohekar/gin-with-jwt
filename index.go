package main

import (
	"golang/gin/config"
	"golang/gin/model"
	"golang/gin/routes"
)

func init() {
	config.ConnDb()
	err := config.DB.AutoMigrate(&model.Candiate{}, &model.Products{}, &model.Delivery{}, &model.Users{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
}

func main() {
	routes.Routers()
}
