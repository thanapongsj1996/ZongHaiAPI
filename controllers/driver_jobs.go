package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"zonghai-api/models"

	"github.com/jinzhu/copier"
)

type DriverJob struct {
	DB *gorm.DB
}

func (d *DriverJob) FindAllDriverJobs(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJobs []models.DriverJob

	if err := d.DB.Preload("Driver").Where(models.DriverJob{IsActive: true}).Find(&driverJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverJobResponseWithDriver
	copier.Copy(&serializedResponse, &driverJobs)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) FindDriverJobByDriverJobUuid(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJobs models.DriverJob
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Preload("Driver").Where("uuid = ?", driverJobUuid).First(&driverJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobResponse
	copier.Copy(&serializedResponse, &driverJobs)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) CreateDriverDeliveryJobResponse(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJob
	driverJobUuid := ctx.Param("driverJobUuid")

	var form createDriverDeliveryJobResponseForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if err := d.DB.Preload("Driver").Where(models.DriverJob{Uuid: driverJobUuid, IsActive: true}).First(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, "Job is not found")
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var jobResponse models.DriverJobDeliveryResponse
	copier.Copy(&jobResponse, &form)
	jobResponse.Uuid = uuid.NewString()
	jobResponse.DriverJob = driverJob

	if err := d.DB.Create(&jobResponse).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse createDriverDeliveryJobResponseResponse
	copier.Copy(&serializedResponse, &jobResponse)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
