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

func (a *Auth) GetDriverProfile(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	sub, _ := ctx.Get("sub")
	driver := sub.(*models.Driver)

	var serializedResponse driverProfile
	copier.Copy(&serializedResponse, driver)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}

func (a *Auth) DriverSignUp(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var form driverAuthForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
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

	var serializedResponse driverAuthResponse
	copier.Copy(&serializedResponse, &driver)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (a *Auth) DriverUpdateProfile(ctx *gin.Context) {
	var jsonResponse models.JSONResponse
	var form driverUpdateProfileForm
	if err := ctx.ShouldBind(&form); err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	sub, _ := ctx.Get("sub")
	driver := sub.(*models.Driver)
	copier.Copy(&driver, &form)

	if err := a.DB.Model(driver).Updates(driver).Error; err != nil {
		errResponse := models.ErrorResponse(jsonResponse, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	var serializedResponse driverUpdateProfileResponse
	copier.Copy(&serializedResponse, driver)

	jsonResponse.Data = serializedResponse
	response := models.SuccessResponse(jsonResponse)
	ctx.JSON(http.StatusOK, response)
}
