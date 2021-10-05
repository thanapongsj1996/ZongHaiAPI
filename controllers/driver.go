package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"zonghai-api/models"
)

type Driver struct {
	DB *gorm.DB
}

func (d *Driver) FindDriverJobs(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJobs []models.DriverJob

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)

	if err := d.DB.Where(models.DriverJob{DriverId: driver.ID}).Find(&driverJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverJobResponse
	copier.Copy(&serializedResponse, &driverJobs)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) CreateDriverJob(ctx *gin.Context) {
	var jsonResponse models.JSONResponse

	var form createDriverJobForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	if !driver.IsVerify || !driver.IsActive {
		errResponse := models.ErrorResponse(jsonResponse, "User is not verify")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	var driverRequest models.DriverJob
	copier.Copy(&driverRequest, &form)
	driverRequest.Uuid = uuid.NewString()
	driverRequest.Driver = driver

	if err := d.DB.Create(&driverRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobResponse
	copier.Copy(&serializedResponse, &driverRequest)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (d *Driver) UpdateDriverJob(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequest models.DriverJob

	var form updateDriverJobForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Where("uuid = ?", driverJobUuid).First(&driverRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if driver.ID != driverRequest.DriverId {
		errResponse := models.ErrorResponse(jsonResponse, "This job is not yours")
		ctx.JSON(http.StatusForbidden, errResponse)
		return
	}

	copier.Copy(&driverRequest, &form)
	if err := d.DB.Model(driverRequest).Updates(&driverRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobResponse
	copier.Copy(&serializedResponse, &driverRequest)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
