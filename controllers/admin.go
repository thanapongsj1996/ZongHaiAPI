package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"os"
	"zonghai-api/models"
)

type Admin struct {
	DB *gorm.DB
}

type deliveryJobResponse struct {
	Uuid           string `json:"uuid"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Items          string `json:"items"`
	Description    string `json:"description"`
	SenderPhone    string `json:"senderPhone"`
	ReceiverPhone  string `json:"receiverPhone"`
	PickupPlace    string `json:"pickupPlace"`
	DeliverPlace   string `json:"deliverPlace"`
	IsDriverAccept bool   `json:"isDriverAccept"`
	DriverJob      struct {
		Uuid string `json:"driverJobUuid"`
	} `json:"driverJob"`
}

type preOrderResponse struct {
	Uuid              string `json:"uuid"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Items             string `json:"items"`
	Description       string `json:"description"`
	Phone             string `json:"phone"`
	DeliverPlace      string `json:"deliverPlace"`
	IsDriverAccept    bool   `json:"isDriverAccept"`
	DriverJobPreOrder struct {
		Uuid string `json:"preOrderJobUuid"`
	} `json:"driverJobPreOrder"`
}

type ProvidedJob struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Description      string `json:"description"`
	Phone            string `json:"phone"`
	DeparturePlace   string `json:"departurePlace"`
	DestinationPlace string `json:"destinationPlace"`
	PlaceOnTheWay    string `json:"placeOnTheWay"`
}

type adminSecret struct {
	Password string `json:"password"`
}

func (a *Admin) GetAllDeliveryResponse(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var jobResponses []models.DriverJobDeliveryResponse

	var form adminSecret

	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if form.Password != os.Getenv("ADMIN_SECRET") {
		errResponse := models.ErrorResponse(jsonResponse, "Incorrect username or password")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := a.DB.Preload("DriverJob").Model(models.DriverJobDeliveryResponse{}).Find(&jobResponses).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []deliveryJobResponse
	copier.Copy(&serializedResponse, &jobResponses)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (a *Admin) GetAllPreOrderResponse(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var jobResponses []models.DriverJobPreOrderResponse

	var form adminSecret

	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if form.Password != os.Getenv("ADMIN_SECRET") {
		errResponse := models.ErrorResponse(jsonResponse, "Incorrect username or password")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := a.DB.Preload("DriverJobPreOrder").Model(models.DriverJobPreOrderResponse{}).Find(&jobResponses).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedResponse []preOrderResponse
	copier.Copy(&serializedResponse, &jobResponses)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

type ProvidedJobsBody struct {
	ProvidedJobs []ProvidedJob `json:"providedJobs"`
	adminSecret
}

func (a *Admin) AddProvidedJobs(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var form ProvidedJobsBody

	if err := ctx.ShouldBindJSON(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	if form.Password != os.Getenv("ADMIN_SECRET") {
		errResponse := models.ErrorResponse(jsonResponse, "Incorrect username or password")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	var providedJobs []models.ProvidedJob
	copier.Copy(&providedJobs, &form.ProvidedJobs)

	for i := range providedJobs {
		providedJobs[i].Uuid = uuid.NewString()
	}

	if err := a.DB.Create(&providedJobs).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
