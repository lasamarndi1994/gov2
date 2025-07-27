package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lasamarndi1994/gov2/api/handler/request"
	"github.com/lasamarndi1994/gov2/helper"
	"github.com/lasamarndi1994/gov2/internal/config"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/internal/mail"
	"github.com/lasamarndi1994/gov2/models"
	"gorm.io/gorm"

	"github.com/lasamarndi1994/gov2/utility/response"
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
	}

	isUnique, field := isEmailOrMobileExists(input.Email, input.MobileNumber)

	if isUnique {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": map[string]string{
				field: field + " already exists",
			}, // convert error to string
		})
		return
	}

	user := models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		MobileNumber: input.MobileNumber,
		Password:     helper.HashPassword(input.Password),
	}

	result := database.DB.Create(&user) // insert data

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  result.Error.Error(), // convert error to string
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   user,
	})
}

func HandleForgotPassword(c *gin.Context) {
	type ForgotRequest struct {
		Email string `json:"email" binding:"required,email"`
	}
	var input ForgotRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	var user models.User
	err := database.DB.Where("email = ?", input.Email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, response.ErrorMessage("email", "Email is not present"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMessage("User registered successfully"))

	go mail.SendHTMLEmail(user.Email, "Welcome !", mail.EmailData{
		Name:  user.FirstName,
		Email: user.Email,
	})
}

func HandleResetPassword(c *gin.Context) {
}

func isEmailOrMobileExists(email string, mobile_number int) (bool, string) {
	var user models.User // 🔥 local variable, not global
	err := database.DB.Where("email = ? OR mobile_number =?", email, mobile_number).First(&user).Error

	if err == nil {
		// Email or mobile  exists → not unique
		if email == user.Email {
			return true, "email"
		}
		if mobile_number == user.MobileNumber {
			return true, "mobile_number"
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Email does NOT exist → unique
		return false, ""
	}
	// DB error (not found and not nil)
	return true, ""
}
