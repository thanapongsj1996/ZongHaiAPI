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

func (c *CustomerJob) FindAllCustomerRequests(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var customerRequests []models.CustomerRequest

	if err := c.DB.Find(&customerRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []customerRequestResponse
	copier.Copy(&serializedResponse, &customerRequests)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *CustomerJob) CreateCustomerRequest(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var form createCustomerRequestForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var customerRequest models.CustomerRequest
	copier.Copy(&customerRequest, &form)

	if err := c.DB.Create(&customerRequest).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse customerRequestResponse
	copier.Copy(&serializedResponse, &customerRequest)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}
