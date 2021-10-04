package routes

import (
	"github.com/gin-gonic/gin"
	"zonghai-api/config"
	"zonghai-api/controllers"
	"zonghai-api/middleware"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")

	v1.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	authenticate := middleware.Authenticate().MiddlewareFunc()

	authGroup := v1.Group("auth")
	authController := controllers.Auth{DB: db}
	{
		authGroup.POST("/driver/sign-up", authController.DriverSignUp)
		authGroup.POST("/driver/sign-in", middleware.Authenticate().LoginHandler)
		authGroup.GET("/driver/profile", authenticate, authController.GetDriverProfile)
		authGroup.PATCH("/driver/profile", authenticate, authController.DriverUpdateProfile)
	}

	customerController := controllers.Customer{DB: db}
	jobsGroup := v1.Group("jobs")
	{
		jobsGroup.GET("/customer-requests", customerController.FindAllCustomerRequests)
		jobsGroup.POST("/customer-requests", customerController.CreateCustomerRequest)
		jobsGroup.DELETE("/customer-requests/all", customerController.ClearCustomerRequest)
	}

}
