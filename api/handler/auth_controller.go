package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/helper"
	"github.com/lasamarndi1994/gov2/internal/config"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/internal/mail"
	"github.com/lasamarndi1994/gov2/internal/request"
	"github.com/lasamarndi1994/gov2/utility/response"
	"github.com/lasamarndi1994/gov2/utility/validation"
	"gorm.io/gorm"
)

var cfg = config.LoadConfig()
var jwtSecret = []byte(cfg.JWTSecretKey)

func HandleLogin(c *gin.Context) {
	var input request.LoginReuest
	var user models.User
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
		c.JSON(http.StatusUnauthorized, response.ErrorMessage("email", "Email is not register"))
		return
	}
	//check password
	if !helper.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, response.ErrorMessage("password", "Enter password is invalid"))
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
		c.JSON(http.StatusUnauthorized, response.ErrorMessage(field, field+" is already register"))
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
	token := helper.GenerateToken(64)

	resetPassword := models.PasswordReset{
		UserId:    user.Id,
		Token:     token,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	result := database.DB.Create(&resetPassword) // insert data

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  result.Error.Error(), // convert error to string
		})
		return
	}

	go mail.SendHTMLEmail(user.Email, "Welcome !", mail.EmailData{
		Name:       user.FirstName,
		Email:      user.Email,
		ResetToken: "https://yourdomain.com/reset-password?token=" + token,
	})

	c.JSON(http.StatusOK, response.SuccessMessage("Sent a link to register email", nil))

}

func HandleResetPassword(c *gin.Context) {

	var input request.ResetPasswordRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	request.InitValidator()
	if errs := request.Validate.Struct(input); errs != nil {
		errs := validation.FormatValidationError(errs)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}
	var resetPasword models.PasswordReset

	result := database.DB.Where("token =? AND expires_at >?", input.Token, time.Now()).First(&resetPasword)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.ErrorMessage("token", "The link is expire"))
		return
	}
	database.DB.Where("token =? ", input.Token).Delete(&resetPasword)

	database.DB.Where("id =? ", resetPasword.UserId).Updates(models.User{Password: input.Password})
	c.JSON(http.StatusOK, response.SuccessMessage("Your password is successfully updated", nil))

}

func isEmailOrMobileExists(email string, mobile_number int) (bool, string) {
	var user models.User //  local variable, not global
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
