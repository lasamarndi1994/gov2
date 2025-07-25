package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lasamarndi1994/gov2/api/handler/request"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/models"

	"github.com/lasamarndi1994/gov2/utility/validation"
)

var jwtSecret = []byte("your-secret-key")

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

	var user models.User
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
