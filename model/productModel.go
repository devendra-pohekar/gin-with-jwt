package model

import (
	"golang/gin/config"
	"golang/gin/response"
)

func AddProduct(data response.AddProductStruct, user_id int) interface{} {
	var table Products

	var product_data = data

	table.ProductName = product_data.ProductName
	table.ID = user_id

	result := config.DB.Create(&table)
	if result.Error != nil {
		return 0
	}
	return 1
}

func FetchProducts(extra_condition int) []Products {
	var product_struct []Products
	db := config.DB
	query := db
	if extra_condition != 0 {
		query = query.Where("ID = ?", extra_condition).First(&product_struct)
	} else {
		query = query.Find(&product_struct)
	}
	result := query
	if result.Error != nil {
		panic("some error occured during the query execution")
	}
	return product_struct
}
