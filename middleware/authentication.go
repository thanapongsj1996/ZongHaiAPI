package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"
	"zonghai-api/config"
	"zonghai-api/models"
)

type loginForm struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

var identityKey = "sub"

func Authenticate() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		// secret key
		Key: []byte(os.Getenv("SECRET_KEY")),

		Timeout: 24 * 90 * time.Hour,

		IdentityKey: identityKey,

		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",

		IdentityHandler: func(c *gin.Context) interface{} {
			var driver models.Driver
			claims := jwt.ExtractClaims(c)
			uuid := claims[identityKey]

			db := config.GetDB()
			if err := db.First(&driver, "uuid = ?", uuid.(string)).Error; err != nil {
				return nil
			}

			return &driver
		},
		// login => driver
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var form loginForm
			var driver models.Driver

			if err := c.ShouldBindJSON(&form); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			db := config.GetDB()
			if err := db.Where("phone = ?", strings.TrimSpace(form.Phone)).First(&driver).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			if err := bcrypt.CompareHashAndPassword([]byte(driver.Password), []byte(form.Password)); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &driver, nil
		},

		// user => payload(sub) => JWT
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Driver); ok {
				claims := jwt.MapClaims{
					identityKey: v.Uuid,
				}

				return claims
			}

			return jwt.MapClaims{}
		},

		// handle error
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error": message,
			})
		},
	})

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	return authMiddleware
}
