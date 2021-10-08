package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	var driverJob models.DriverJob
	copier.Copy(&driverJob, &form)
	driverJob.Uuid = uuid.NewString()
	driverJob.Driver = driver

	if err := d.DB.Create(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobResponse
	copier.Copy(&serializedResponse, &driverJob)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (d *Driver) UpdateDriverJob(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJob

	var form updateDriverJobForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Where("uuid = ?", driverJobUuid).First(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if driver.ID != driverJob.DriverId {
		errResponse := models.ErrorResponse(jsonResponse, "This job is not yours")
		ctx.JSON(http.StatusForbidden, errResponse)
		return
	}

	copier.Copy(&driverJob, &form)
	if err := d.DB.Model(driverJob).Updates(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobResponse
	copier.Copy(&serializedResponse, &driverJob)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) FindDriverJobsDetail(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJob

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Preload("DriverJobDeliveryResponses").Where(models.DriverJob{DriverId: driver.ID, Uuid: driverJobUuid}).First(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobResponseWithResponses
	copier.Copy(&serializedResponse, &driverJob)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) SetDeliveryJobIsAcceptResponse(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJob

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverJobUuid := ctx.Param("driverJobUuid")
	responseUuid := ctx.Param("responseUuid")
	acceptValue, _ := strconv.ParseBool(ctx.Param("acceptValue"))

	if err := d.DB.Where(models.DriverJob{DriverId: driver.ID, Uuid: driverJobUuid}).Find(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var jobResponse models.DriverJobDeliveryResponse

	if err := d.DB.Model(&jobResponse).Where(models.DriverJobDeliveryResponse{Uuid: responseUuid}).Update("is_driver_accept", acceptValue).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) FindDriverJobsPreOrder(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJobs []models.DriverJobPreOrder

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)

	if err := d.DB.Where(models.DriverJobPreOrder{DriverId: driver.ID}).Find(&driverJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverJobPreOrderResponse
	copier.Copy(&serializedResponse, &driverJobs)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) FindDriverJobsPreOrderDetail(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJobPreOrder

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Preload("DriverJobPreOrderResponses").Where(models.DriverJobPreOrder{DriverId: driver.ID, Uuid: driverJobUuid}).First(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobPreOrderResponseWithResponses
	copier.Copy(&serializedResponse, &driverJob)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *Driver) CreateDriverJobPreOrder(ctx *gin.Context) {
	var jsonResponse models.JSONResponse

	var form createDriverJobPreOrderForm
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

	var driverJobPreOrder models.DriverJobPreOrder
	copier.Copy(&driverJobPreOrder, &form)
	driverJobPreOrder.Uuid = uuid.NewString()
	driverJobPreOrder.Driver = driver

	if err := d.DB.Create(&driverJobPreOrder).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobPreOrderResponse
	copier.Copy(&serializedResponse, &driverJobPreOrder)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (d *Driver) SetPreOrderJobIsAcceptResponse(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJobPreOrder

	driverSub, _ := ctx.Get("sub")
	driver := *driverSub.(*models.Driver)
	driverJobUuid := ctx.Param("driverJobUuid")
	responseUuid := ctx.Param("responseUuid")
	acceptValue, _ := strconv.ParseBool(ctx.Param("acceptValue"))

	if err := d.DB.Where(models.DriverJobPreOrder{DriverId: driver.ID, Uuid: driverJobUuid}).Find(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var jobResponse models.DriverJobPreOrderResponse

	if err := d.DB.Model(&jobResponse).Where(models.DriverJobPreOrderResponse{Uuid: responseUuid}).Update("is_driver_accept", acceptValue).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
