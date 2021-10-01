package routes

import (
	"github.com/gin-gonic/gin"
	"zonghai-api/config"
	"zonghai-api/controllers"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")

	v1.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	customerController := controllers.Customer{DB: db}
	jobsGroup := v1.Group("jobs")
	{
		jobsGroup.GET("/customer-requests", customerController.FindAllCustomerRequests)
		jobsGroup.POST("/customer-requests", customerController.CreateCustomerRequest)
		jobsGroup.DELETE("/customer-requests/all", customerController.ClearCustomerRequest)
	}

}
