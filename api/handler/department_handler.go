package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/utility/response"
	"github.com/lasamarndi1994/gov2/utility/validation"
)

type DepartmentRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func GetDepartments(c *gin.Context) {
}

func CreateDepartment(c *gin.Context) {
	var input DepartmentRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		if errs != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"erros": errs})
		}
	}
	department := models.Department{
		Name:   input.Name,
		Status: input.Status,
	}

	result := database.DB.Create(&department)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage("name", result.Error.Error()))
		return
	}

	c.JSON(http.StatusAccepted, response.ResponseData("Success", department))
}
func ShowDepartment(c *gin.Context) {
}

func DeleteDepartment(c *gin.Context) {
}
