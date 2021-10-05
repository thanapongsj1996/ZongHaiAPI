package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"zonghai-api/models"

	"github.com/jinzhu/copier"
)

type CustomerJob struct {
	DB *gorm.DB
}

func (c *CustomerJob) FindAllCustomerJobs(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var customerJobs []models.CustomerJob

	if err := c.DB.Find(&customerJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []customerJobResponse
	copier.Copy(&serializedResponse, &customerJobs)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *CustomerJob) CreateCustomerJob(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var form createCustomerJobForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var customerJob models.CustomerJob
	copier.Copy(&customerJob, &form)

	if err := c.DB.Create(&customerJob).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse customerJobResponse
	copier.Copy(&serializedResponse, &customerJob)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}
