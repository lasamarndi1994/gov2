package handler

import (
	"net/http"
	"time"

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
	joinngDate, _ := time.Parse("2006-01-02", input.JoiningDate)
	dateofBirth, _ := time.Parse("2006-01-02", input.DateofBirth)

	user := models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		MobileNumber: input.MobileNumber,
		EmployeeId:   input.EmployeeId,
		AboutMe:      input.AboutMe,
		JoiningDate:  joinngDate,
		DateofBirth:  dateofBirth,
		UserDesignations: []models.UserDesignation{
			{DesiginationId: input.DesignationId},
		},
		UserDepartments: []models.UserDepartment{
			{DepartmentId: input.DepartmentId},
		},
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadGateway, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseData("Success", user))
}
