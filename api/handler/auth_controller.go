package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lasamarndi1994/gov2/api/handler/request"
	"github.com/lasamarndi1994/gov2/helper"
	"github.com/lasamarndi1994/gov2/internal/config"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/models"

	"github.com/lasamarndi1994/gov2/utility/validation"
)

var cfg = config.LoadConfig()
var jwtSecret = []byte(cfg.JWTSecretKey)
var user models.User

func HandleLogin(c *gin.Context) {
	var input request.LoginReuest

	if err := c.ShouldBindJSON(&input); err != nil {
		// Validation errors
		errs := validation.FormatValidationError(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	//check email

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "The user is not register",
		})
		return
	}
	//check password
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"token":  tokenString,
		"user":   user,
	})
}

func HandleRegister(c *gin.Context) {
	var input request.RegisterRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	} else {
		usr := models.User{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  helper.HashPassword(input.Password),
		}

		if err := database.DB.Create(&usr); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": usr,
		})
	}

}
