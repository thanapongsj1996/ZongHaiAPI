package controllers

import (
	"github.com/gin-gonic/gin"
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
	var driverRequests []models.DriverJob

	if err := d.DB.Preload("Driver").Where(models.DriverJob{IsActive: true}).Find(&driverRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverJobResponseWithDriver
	copier.Copy(&serializedResponse, &driverRequests)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) FindDriverJobByDriverJobUuid(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequests models.DriverJob
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Preload("Driver").Where("uuid = ?", driverJobUuid).First(&driverRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverJobResponse
	copier.Copy(&serializedResponse, &driverRequests)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
