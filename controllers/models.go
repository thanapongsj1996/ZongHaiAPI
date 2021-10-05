package controllers

import "time"

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

type driverUpdateProfileForm struct {
	FirstName  string `form:"firstName"`
	LastName   string `form:"lastName"`
	Address    string `form:"address"`
	ProfileImg string `form:"profileImg"`
}

type driverUpdateProfileResponse struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Address    string `json:"address"`
	ProfileImg string `json:"profileImg"`
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
	IsActive        bool    `form:"isActive"`
}

type driverRequestResponse struct {
	Uuid             string    `json:"uuid"`
	Description      string    `json:"description"`
	StartPrice       float64   `json:"startPrice"`
	Phone            string    `json:"phone"`
	DeparturePlace   string    `json:"departurePlace"`
	DepartureTime    time.Time `json:"departureTime" time_format:"2006-01-02T15:04:05+0700"`
	DestinationPlace string    `json:"destinationPlace"`
	DestinationTime  time.Time `json:"destinationTime" time_format:"2006-01-02T15:04:05+0700"`
	PlaceOnTheWay    string    `json:"placeOnTheWay"`
	IsActive         bool      `json:"isActive"`
}

type driverRequestResponseWithDriver struct {
	driverRequestResponse
	Driver struct {
		Uuid       string `json:"uuid"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		ProfileImg string `json:"profileImg"`
	} `json:"driver"`
}

type driverRequestByDriverUuidResponse struct {
	driverRequestResponse
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
	IsActive         bool      `form:"isActive"`
}

type updateDriverRequestForm struct {
	createDriverRequestForm
}