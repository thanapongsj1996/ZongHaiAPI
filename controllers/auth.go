package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"zonghai-api/models"
)

type Auth struct {
	DB *gorm.DB
}

type driverAuthForm struct {
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type driverAuthResponse struct {
	Uuid      string `json:"uuid"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type driverProfile struct {
	Uuid             string `json:"uuid"`
	Phone            string `json:"phone"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Address          string `json:"address"`
	DriverLicenseID  string `json:"driverLicenseID"`
	DriverLicenseImg string `json:"driverLicenseImg"`
	ProfileImg       string `json:"profileImg"`
	IsVerify         bool   `json:"isVerify"`
	IsActive         bool   `json:"isActive"`
}

func (a *Auth) GetDriverProfile(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	sub, _ := ctx.Get("sub")
	driver := sub.(*models.Driver)

	var serializedDriver driverProfile
	copier.Copy(&serializedDriver, driver)

	jsonResponse.Data = serializedDriver
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (a *Auth) DriverSignUp(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var form driverAuthForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var driver models.Driver
	copier.Copy(&driver, &form)
	driver.Uuid = uuid.NewString()
	driver.Password = driver.GenerateEncryptedPassword()
	if err := a.DB.Create(&driver).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	var serializedUser driverAuthResponse
	copier.Copy(&serializedUser, &driver)

	jsonResponse.Data = serializedUser
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}
