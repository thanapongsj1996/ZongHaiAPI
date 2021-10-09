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
		driverJobGroup.POST("/:driverJobUuid", driverJobController.CreateDriverDeliveryJobResponse)

		driverJobGroup.GET("/pre-order", driverJobController.FindAllDriverJobsPreOrder)
		driverJobGroup.GET("pre-order/:driverJobUuid", driverJobController.FindDriverJobsPreOrderByJobUuid)
		driverJobGroup.POST("pre-order/:driverJobUuid", driverJobController.CreateDriverPreOrderJobResponse)
	}

	driverController := controllers.Driver{DB: db}
	driverGroup := v1.Group("driver")
	{
		driverGroup.GET("/jobs", authenticate, driverController.FindDriverJobs)
		driverGroup.POST("/jobs", authenticate, driverController.CreateDriverJob)
		driverGroup.PATCH("/jobs/:driverJobUuid", authenticate, driverController.UpdateDriverJob)
		driverGroup.GET("/jobs/:driverJobUuid", authenticate, driverController.FindDriverJobsDetail)
		driverGroup.PATCH("/jobs/:driverJobUuid/:responseUuid/accept/:acceptValue", authenticate, driverController.SetDeliveryJobIsAcceptResponse)

		driverGroup.GET("/jobs/pre-order", authenticate, driverController.FindDriverJobsPreOrder)
		driverGroup.GET("/jobs/pre-order/:driverJobUuid", authenticate, driverController.FindDriverJobsPreOrderDetail)
		driverGroup.POST("/jobs/pre-order", authenticate, driverController.CreateDriverJobPreOrder)
		driverGroup.PATCH("/jobs/pre-order/:driverJobUuid/:responseUuid/accept/:acceptValue", authenticate, driverController.SetPreOrderJobIsAcceptResponse)
	}

	adminController := controllers.Admin{DB: db}
	adminGroup := v1.Group("admin")
	{
		adminGroup.POST("/delivery-jobs/responses", adminController.GetAllDeliveryResponse)
		adminGroup.POST("/preorder-jobs/responses", adminController.GetAllPreOrderResponse)
	}
}
