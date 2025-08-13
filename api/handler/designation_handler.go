package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/utility/response"
	"github.com/lasamarndi1994/gov2/utility/validation"
)

type DesignationRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func GetDesignations(c *gin.Context) {
	var Designations []models.Designation
	database.DB.Order("id desc").Find(&Designations)
	c.JSON(http.StatusOK, response.SuccessResponse(Designations))
}

func CreateDesignation(c *gin.Context) {
	var input DesignationRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		if errs != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"erros": errs})
		}
	}
	Designation := models.Designation{
		Name:   input.Name,
		Status: input.Status,
	}
	err := database.DB.Create(&Designation).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(Designation))
}
func UpdateDesignation(c *gin.Context) {
	id := c.Param("id")
	var Designation models.Designation

	if err := database.DB.First(&Designation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.ErrorMessage("message", "Data not found"))
		return
	}

	var input DesignationRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}
	Designation.Name = input.Name
	Designation.Status = input.Status

	if err := database.DB.Save(&Designation).Error; err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage("name", "The name is alreay taken"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(Designation))
}
func ShowDesignation(c *gin.Context) {
	var id = c.Param("id")
	var Designation models.Designation
	if err := database.DB.First(&Designation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.ErrorMessage("message", "Data not found"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(Designation))

}

func DeleteDesignation(c *gin.Context) {
	var id = c.Param("id")
	var Designation models.Designation
	if err := database.DB.First(&Designation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.ErrorMessage("message", "Data not found"))
		return
	}
	if err := database.DB.Delete(&Designation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessMessage("Successfully Deleted"))
}
