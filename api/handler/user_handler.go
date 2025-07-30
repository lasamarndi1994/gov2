package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/internal/request"
	"github.com/lasamarndi1994/gov2/utility/response"
	"github.com/lasamarndi1994/gov2/utility/validation"
)

func GetUserDetails(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, response.ResponseData("Success", users))
}

func CreateEmployee(c *gin.Context) {

	var input request.EmployeeRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	var user models.User
	err := database.DB.Create(&user).Error
	if err != nil {
		//c.JSON(http.StatusBadGateway, response.ErrorMessage("Something went wrong"))
	}
	c.JSON(http.StatusOK, response.ResponseData("Success", user))
}
