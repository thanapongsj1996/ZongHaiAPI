package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"zonghai-api/models"

	"github.com/jinzhu/copier"
)

type Customer struct {
	DB *gorm.DB
}

type customerRequestResponse struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Phone           string  `json:"phone"`
	PickupLocation  string  `json:"pickupLocation"`
	DeliverLocation string  `json:"deliverLocation"`
}

type createCustomerRequestForm struct {
	Title           string  `form:"title" binding:"required"`
	Description     string  `form:"description"`
	Price           float64 `form:"price" binding:"required"`
	Phone           string  `form:"phone" binding:"required"`
	PickupLocation  string  `form:"pickupLocation" binding:"required"`
	DeliverLocation string  `form:"deliverLocation" binding:"required"`
}

func (c *Customer) FindAllCustomerRequests(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var customerRequests []models.CustomerRequest

	if err := c.DB.Find(&customerRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	var serializedArticles []customerRequestResponse
	copier.Copy(&serializedArticles, &customerRequests)

	jsonResponse.Data = serializedArticles
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *Customer) CreateCustomerRequest(ctx *gin.Context) {
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

	var serializedArticles customerRequestResponse
	copier.Copy(&serializedArticles, &customerRequest)

	jsonResponse.Data = serializedArticles
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}
