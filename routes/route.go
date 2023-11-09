package routes

import (
	"golang/gin/controller"
	middlware "golang/gin/middlwares"

	"github.com/gin-gonic/gin"
)

func Routers() {
	var ApiAuth = middlware.AuthMiddleware()
	router := gin.Default()

	router.POST("/login", middlware.Login)

	router.GET("/", controller.Sample)

	router.POST("/sample", controller.GetAllSample)
	router.POST("/insert_user", ApiAuth, controller.InsertCandidate)
	router.POST("/fetchall_candidates", ApiAuth, controller.FetchALlCandidates)
	router.POST("/fetchby_id", ApiAuth, controller.FetchById)
	router.POST("/filter_candidate", ApiAuth, controller.FilterCandidateByName)
	router.POST("/register_user", controller.RegisterUser)

	router.POST("/add_product", ApiAuth, controller.AddProducts)

	router.Run()
}
