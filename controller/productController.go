package controller

import (
	"golang/gin/helpers"
	"golang/gin/model"
	"golang/gin/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProducts(p *gin.Context) {
	var product_struct response.AddProductStruct
	err := p.ShouldBindJSON(product_struct)
	if err != nil {
		p.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}
	user_id := helpers.GetUserFromToken(p)
	result := model.AddProduct(product_struct, user_id)
	if result != 0 {
		p.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}
	p.JSON(http.StatusAccepted, gin.H{
		"success": "successfully insert product",
	})
}

func FetchProducts(p *gin.Context) {
	var id response.FetchProduct
	err := p.ShouldBindJSON(&id)
	var extra_condition int = 0
	if id.ProductID != 0 {
		extra_condition = id.ProductID
	}

	if err != nil {
		p.JSON(http.StatusInternalServerError, gin.H{
			"error": "Please Provide The Essential Details",
		})
		return
	}
	result := model.FetchProducts(extra_condition)

	log.Print(result)
	p.JSON(200, gin.H{
		"status":  200,
		"Message": "successfully Found Products ",
		"data":    result,
	})
}
