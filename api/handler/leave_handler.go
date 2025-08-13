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

func ApplyLeave(c *gin.Context) {
	var input request.Leave

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	var leave = models.Leave{
		LeaveTypeId:       input.LeaveTypeId,
		FromDate:          input.FromDate,
		ToDate:            input.ToDate,
		FromLeaveValue:    input.FromLeaveValue,
		ToLeaveValue:      input.ToLeaveValue,
		Remarks:           input.Remarks,
		EmailNotification: input.EmailNotification,
	}

	err := database.DB.Create(&leave).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusAccepted, response.SuccessResponse(leave))
}

func UpdateLeave(c *gin.Context) {
}

func DeleteLeave(c *gin.Context) {
}

func CreateLeave(c *gin.Context) {
	var input request.LevaeTypeRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errors := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}

	var leave = models.LevaeType{
		Name:   input.Name,
		Status: input.Status,
	}
	err := database.DB.Create(&leave).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(leave))

}
