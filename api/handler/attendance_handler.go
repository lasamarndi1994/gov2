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

func GetAttendance(c *gin.Context) {
	var attendances []models.Attendance
	c.JSON(http.StatusOK, response.SuccessResponse(&attendances))
}

func StoreAttendance(c *gin.Context) {
	var input request.AttendanceRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	attendance := models.Attendance{
		UserId:         input.UserId,
		AttendanceDate: input.AttendanceDate,
		CheckIn:        input.CheckIn,
		CheckOut:       input.CheckOut,
	}

	err := database.DB.Create(&attendance).Error
	if err != nil {
		c.JSON(http.StatusBadGateway, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(&attendance))
}

func UpdateAttendance(c *gin.Context) {

}
