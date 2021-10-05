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

	customerRequestController := controllers.CustomerJob{DB: db}
	customerRequestGroup := v1.Group("customer-request")
	{
		customerRequestGroup.GET("", customerRequestController.FindAllCustomerRequests)
		customerRequestGroup.POST("", customerRequestController.CreateCustomerRequest)
	}

	driverRequestController := controllers.DriverJob{DB: db}
	driverRequestGroup := v1.Group("driver-request")
	{
		driverRequestGroup.GET("", driverRequestController.FindAllDriverRequests)
		driverRequestGroup.GET("/:driverRequestUuid", driverRequestController.FindDriverRequestByDriverRequestUuid)
	}

	driverController := controllers.Driver{DB: db}
	driverGroup := v1.Group("driver")
	{
		driverGroup.GET("/jobs", authenticate, driverController.FindDriverRequests)
		driverGroup.POST("/jobs", authenticate, driverController.CreateDriverRequest)
		driverGroup.PATCH("/jobs/:driverRequestUuid", authenticate, driverController.UpdateDriverRequest)
	}
}
