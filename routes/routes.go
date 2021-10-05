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

	customerJobController := controllers.CustomerJob{DB: db}
	customerJobGroup := v1.Group("customer-jobs")
	{
		customerJobGroup.GET("", customerJobController.FindAllCustomerJobs)
		customerJobGroup.POST("", customerJobController.CreateCustomerJob)
	}

	driverJobController := controllers.DriverJob{DB: db}
	driverJobGroup := v1.Group("driver-jobs")
	{
		driverJobGroup.GET("", driverJobController.FindAllDriverJobs)
		driverJobGroup.GET("/:driverJobUuid", driverJobController.FindDriverJobByDriverJobUuid)
	}

	driverController := controllers.Driver{DB: db}
	driverGroup := v1.Group("driver")
	{
		driverGroup.GET("/jobs", authenticate, driverController.FindDriverJobs)
		driverGroup.POST("/jobs", authenticate, driverController.CreateDriverJob)
		driverGroup.PATCH("/jobs/:driverJobUuid", authenticate, driverController.UpdateDriverJob)
	}
}
