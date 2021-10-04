package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zonghai-api/config"
	"zonghai-api/controllers"
	"zonghai-api/middleware"
	"zonghai-api/models"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")

	v1.GET("/health-check", func(ctx *gin.Context) {
		var jsonResponse models.JSONResponse
		response := models.SuccessResponse(jsonResponse)
		ctx.JSON(http.StatusOK, response)
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
	driverController := controllers.Driver{DB: db}
	jobsGroup := v1.Group("jobs")
	{
		// Customer Request Endpoints
		jobsGroup.GET("/customer", customerController.FindAllCustomerRequests)
		jobsGroup.POST("/customer", customerController.CreateCustomerRequest)

		// Driver Request Endpoints
		jobsGroup.GET("/driver", driverController.FindAllDriverRequests)
		jobsGroup.POST("/driver", driverController.CreateDriverRequest)
	}

}
