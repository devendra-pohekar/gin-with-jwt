package model

import (
	"golang/gin/config"
	"golang/gin/helpers"
	"golang/gin/response"
	"time"
)

func AddCandidate(data response.CandidateInsert) interface{} {
	var candidate = data
	var table Candiate
	hashpass, _ := helpers.HashPassword(candidate.Password)
	table.Email = candidate.Email
	table.Name = candidate.Name
	table.UserName = candidate.UserName
	table.Password = hashpass
	table.CreatedAt = time.Now()
	result := config.DB.Create(&table)
	if result.Error != nil {
		panic("Faild to insert data : " + result.Error.Error())
	}
	return result

}

func FetchALlCandidates() []Candiate {
	var candidates []Candiate

	result := config.DB.Find(&candidates)
	if result.Error != nil {
		panic("Faild to find data : " + result.Error.Error())
	}
	return candidates
}

func FetchCandidateById(extra_condition int) []Candiate {
	var candidate_struct []Candiate
	db := config.DB
	query := db
	if extra_condition != 0 {
		query = query.Where("ID = ?", extra_condition).First(&candidate_struct)
	} else {
		query = query.Find(&candidate_struct)
	}
	result := query
	if result.Error != nil {
		panic("some error occured during the query execution")
	}
	return candidate_struct
}
func FilterCandidateByName(filter_string string) []Candiate {
	var candidate_struct []Candiate
	result := config.DB.Where("ID LIKE ? OR Name LIKE ?", filter_string, filter_string).Find(&candidate_struct)
	if result.Error != nil {
		panic("some error occurred in query")
	}
	return candidate_struct
}

func LoginUser(email string, password string) Users {
	var user Users

	result := config.DB.Where("Email = ? AND Password = ?", email, password).Find(&user)
	if result.Error != nil {
		panic("cannot find user data")
	}
	return user

}

func RegisterUser(data response.UserRegister) interface{} {
	var user_data = data
	var table Users
	table.Email = user_data.Email
	table.Name = user_data.Name
	table.Password = user_data.Password
	table.CreatedAt = time.Now()
	result := config.DB.Create(&table)
	return result

}
