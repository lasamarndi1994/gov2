package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/internal/request"
	"github.com/lasamarndi1994/gov2/utility/response"
	"github.com/lasamarndi1994/gov2/utility/validation"
)

func ApplyLeave(c *gin.Context) {
	var input request.LeaveReuest

	if err := c.ShouldBindWith(&input, binding.FormMultipart); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}
	userIDValue, _ := c.Get("userID")
	form_date, _ := time.Parse("2006-01-02", input.FromDate)
	to_date, _ := time.Parse("2006-01-02", input.ToDate)

	jsonStr, _ := json.Marshal(input.EmailNotification)

	if input.Attachement != nil {
		c.SaveUploadedFile(input.Attachement, "uploads/leave/"+input.Attachement.Filename)
	}

	var leave = models.Leave{
		LeaveTypeId:       input.LeaveTypeId,
		UserId:            userIDValue.(uint),
		FromDate:          form_date,
		ToDate:            to_date,
		FromLeaveValue:    input.FromLeaveValue,
		ToLeaveValue:      input.ToLeaveValue,
		Remarks:           input.Remarks,
		Attachement:       input.Attachement.Filename,
		EmailNotification: string(jsonStr),
		LeaveStatus:       "Pending",
	}

	err := database.DB.Create(&leave).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusAccepted, response.SuccessResponse(leave))
}

func UpdateLeave(c *gin.Context) {
	var input request.LeaveReuest
	if err := c.ShouldBindWith(&input, binding.FormMultipart); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}
	id := c.Param("id")

	var leave models.Leave
	if err := database.DB.First(&leave, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}

	form_date, _ := time.Parse("2006-01-02", input.FromDate)
	to_date, _ := time.Parse("2006-01-02", input.ToDate)
	jsonStr, _ := json.Marshal(input.EmailNotification)

	if input.Attachement != nil {
		c.SaveUploadedFile(input.Attachement, "uploads/leave/"+input.Attachement.Filename)
	}

	leave.LeaveTypeId = input.LeaveTypeId
	leave.FromDate = form_date
	leave.ToDate = to_date
	leave.FromLeaveValue = input.FromLeaveValue
	leave.ToLeaveValue = input.ToLeaveValue
	leave.Remarks = input.Remarks
	leave.Attachement = input.Attachement.Filename
	leave.EmailNotification = string(jsonStr)
	leave.LeaveStatus = input.LeaveStatus

	if err := database.DB.Save(&leave).Error; err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMessage("Successfully updated."))
}

func DeleteLeave(c *gin.Context) {
	id := c.Param("id")
	var leave models.Leave

	err := database.DB.Where("id =?", id).Delete(&leave).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMessage("Successfully deleted"))
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
