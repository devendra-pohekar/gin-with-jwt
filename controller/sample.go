package controller

import (
	"golang/gin/model"
	"golang/gin/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sample(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

type SampleStruct struct {
	SampleName string
	City       string
	Age        int
}

type Candidate struct {
	Name     string `json:"candidate_name" `
	Email    string `json:"candidate_email" `
	UserName string `json:"candidate_username" `
}

type CandidateId struct {
	ID int `json:"candidate_id"`
}

func GetAllSample(c *gin.Context) {
	result := SampleStruct{Age: 20, City: "Burhanpur Madhya Pradesh", SampleName: "hello"}
	c.JSON(200, gin.H{
		"response": result,
	})
}

// func InsertCandidate(c *gin.Context) {
// 	var candidate model.Candiate
// 	err := c.ShouldBindJSON(&candidate)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Please Provide The Essential Details",
// 		})
// 		return
// 	}

// 	result := model.AddCandidate(candidate)
// 	if result != 0 {
// 		c.JSON(http.StatusAlreadyReported, gin.H{
// 			"result": "successfully inserted candidate",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"error": "something wrong to insert candidate",
// 	})

// }

func RegisterUser(c *gin.Context) {
	var user response.UserRegister
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := model.RegisterUser(user)
	if result != 0 {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"result": "successfully Register User",
		})
		return
	}
}

func InsertCandidate(c *gin.Context) {
	var candidate response.CandidateInsert
	err := c.ShouldBindJSON(&candidate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Please Provide The Essential Details",
		})
		return
	}

	result := model.AddCandidate(candidate)
	if result != 0 {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"result": "successfully inserted candidate",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "something wrong to insert candidate",
	})

}

func FetchALlCandidates(c *gin.Context) {
	result := model.FetchALlCandidates()
	c.JSON(200, gin.H{
		"status":  200,
		"Message": "successfully Found Candidates ",
		"data":    result,
	})
}

func FetchById(c *gin.Context) {
	var id CandidateId
	err := c.ShouldBindJSON(&id)
	var extra_condition int = 0
	if id.ID != 0 {
		extra_condition = id.ID
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Please Provide The Essential Details",
		})
		return
	}
	result := model.FetchCandidateById(extra_condition)

	log.Print(result)
	c.JSON(200, gin.H{
		"status":  200,
		"Message": "successfully Found Candidates ",
		"data":    result,
	})
}

func FilterCandidateByName(c *gin.Context) {
	var filter_string response.FilterStruct
	err := c.ShouldBindJSON(&filter_string)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Please Provide The Essential Details",
		})
		return
	}
	log.Print(filter_string)
	result := model.FilterCandidateByName(filter_string.Name)
	c.JSON(http.StatusFound, gin.H{
		"message": "successfully found data",
		"data":    result,
	})
}
