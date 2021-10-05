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

func (d *DriverJob) FindAllDriverRequests(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequests []models.DriverRequest

	if err := d.DB.Preload("Driver").Where(models.DriverRequest{IsActive: true}).Find(&driverRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverRequestResponseWithDriver
	copier.Copy(&serializedResponse, &driverRequests)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) FindDriverRequestByDriverRequestUuid(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequests models.DriverRequest
	driverRequestUuid := ctx.Param("driverRequestUuid")

	if err := d.DB.Preload("Driver").Where("uuid = ?", driverRequestUuid).First(&driverRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverRequestResponse
	copier.Copy(&serializedResponse, &driverRequests)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
