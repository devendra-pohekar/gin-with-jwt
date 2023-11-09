package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Dev@123"
	dbname   = "golang_practice"
)

func ConnDb() {

	credential_details := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(credential_details), &gorm.Config{})
	if err != nil {
		panic("Failed To Connect DB: " + err.Error())
	}
	DB = db
	fmt.Println("successfully connected db")

}
