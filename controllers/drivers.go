package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
	"zonghai-api/models"

	"github.com/jinzhu/copier"
)

type Driver struct {
	DB *gorm.DB
}

type driverRequestResponse struct {
	Description      string    `json:"description"`
	StartPrice       float64   `json:"startPrice"`
	Phone            string    `json:"phone"`
	DeparturePlace   string    `json:"departurePlace"`
	DepartureTime    time.Time `json:"departureTime" time_format:"2006-01-02T15:04:05+0700"`
	DestinationPlace string    `json:"destinationPlace"`
	DestinationTime  time.Time `json:"destinationTime" time_format:"2006-01-02T15:04:05+0700"`
	PlaceOnTheWay    string    `json:"placeOnTheWay"`
}

type createDriverRequestForm struct {
	Description      string    `form:"description"`
	StartPrice       float64   `form:"startPrice"`
	Phone            string    `form:"phone"`
	DeparturePlace   string    `form:"departurePlace"`
	DepartureTime    time.Time `form:"departureTime" time_format:"2006-01-02T15:04:05+0700"`
	DestinationPlace string    `form:"destinationPlace"`
	DestinationTime  time.Time `form:"destinationTime" time_format:"2006-01-02T15:04:05+0700"`
	PlaceOnTheWay    string    `form:"placeOnTheWay"`
}

func (d *Driver) FindAllDriverRequests(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var driverRequests []models.DriverRequest

	if err := d.DB.Find(&driverRequests).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
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

	var driverRequest models.DriverRequest
	copier.Copy(&driverRequest, &form)
	driverRequest.Uuid = uuid.NewString()

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
