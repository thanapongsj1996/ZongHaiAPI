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

func (d *Driver) FindDriverRequests(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequests []models.DriverRequest

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)

	if err := d.DB.Where(models.DriverRequest{DriverId: driver.ID}).Find(&driverRequests).Error; err != nil {
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

func (d *Driver) CreateDriverRequest(ctx *gin.Context) {
	var jsonResponse models.JSONResponse

	var form createDriverRequestForm
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

	var driverRequest models.DriverRequest
	copier.Copy(&driverRequest, &form)
	driverRequest.Uuid = uuid.NewString()
	driverRequest.Driver = driver

	if err := d.DB.Create(&driverRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverRequestResponse
	copier.Copy(&serializedResponse, &driverRequest)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (d *Driver) UpdateDriverRequest(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequest models.DriverRequest

	var form updateDriverRequestForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverRequestUuid := ctx.Param("driverRequestUuid")

	if err := d.DB.Where("uuid = ?", driverRequestUuid).First(&driverRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if driver.ID != driverRequest.DriverId {
		errResponse := models.ErrorResponse(jsonResponse, "This request is not yours")
		ctx.JSON(http.StatusForbidden, errResponse)
		return
	}

	copier.Copy(&driverRequest, &form)
	if err := d.DB.Model(driverRequest).Updates(&driverRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverRequestResponse
	copier.Copy(&serializedResponse, &driverRequest)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) FindDriverRequestsByDriverUuid(ctx *gin.Context) {
	var jsonResponse models.JSONResponse

	var driver models.Driver
	var driverRequests []models.DriverRequest
	driverUuid := ctx.Param("driverUuid")

	if err := d.DB.Where("uuid = ?", driverUuid).First(&driver).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if err := d.DB.Where("driver_id = ?", driver.ID).Find(&driverRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverRequestByDriverUuidResponse
	copier.Copy(&serializedResponse, &driverRequests)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
