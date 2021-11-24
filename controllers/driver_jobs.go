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

	var serializedResponse driverJobResponseWithDriver
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
	ctx.JSON(http.StatusCreated, response)
}

func (d *DriverJob) FindAllDriverJobsPreOrder(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJobsPreOrder []models.DriverJobPreOrder

	if err := d.DB.Preload("Driver").Where(models.DriverJobPreOrder{IsActive: true}).Find(&driverJobsPreOrder).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []driverJobPreOrderResponseWithDriver
	copier.Copy(&serializedResponse, &driverJobsPreOrder)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) FindDriverJobsPreOrderByJobUuid(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJobsPreOrder models.DriverJobPreOrder
	driverJobUuid := ctx.Param("driverJobUuid")

	if err := d.DB.Preload("Driver").Where(models.DriverJobPreOrder{IsActive: true, Uuid: driverJobUuid}).First(&driverJobsPreOrder).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse driverJobPreOrderResponseWithDriver
	copier.Copy(&serializedResponse, &driverJobsPreOrder)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) CreateDriverPreOrderJobResponse(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverJob models.DriverJobPreOrder
	driverJobUuid := ctx.Param("driverJobUuid")

	var form createDriverPreOrderJobResponseForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if err := d.DB.Preload("Driver").Where(models.DriverJobPreOrder{Uuid: driverJobUuid, IsActive: true}).First(&driverJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, "Job is not found")
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var jobResponse models.DriverJobPreOrderResponse
	copier.Copy(&jobResponse, &form)
	jobResponse.Uuid = uuid.NewString()
	jobResponse.DriverJobPreOrder = driverJob

	if err := d.DB.Create(&jobResponse).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse createDriverPreOrderJobResponseResponse
	copier.Copy(&serializedResponse, &jobResponse)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}

type ProvidedJobsResponse struct {
	Uuid             string  `json:"uuid"`
	FirstName        string  `json:"firstName"`
	LastName         string  `json:"lastName"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	Phone            string  `json:"phone"`
	DeparturePlace   string  `json:"departurePlace"`
	DestinationPlace string  `json:"destinationPlace"`
	PlaceOnTheWay    string  `json:"placeOnTheWay"`
	IsActive         bool    `json:"isActive"`
}

func (d *DriverJob) FindProvidedJobs(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var providedJobs []models.ProvidedJob

	if err := d.DB.Find(&providedJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []ProvidedJobsResponse
	copier.Copy(&serializedResponse, &providedJobs)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (d *DriverJob) FindProvidedJobsByUuid(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var providedJob models.ProvidedJob

	providedJobUuid := ctx.Param("providedJobUuid")
	if err := d.DB.Where(models.ProvidedJob{Uuid: providedJobUuid}).Find(&providedJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse ProvidedJobsResponse
	copier.Copy(&serializedResponse, &providedJob)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
