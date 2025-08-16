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

func GetHolidays(c *gin.Context) {
	var holidys []models.Holiday
	database.DB.Order("id desc").Find(&holidys)
	c.JSON(http.StatusOK, response.SuccessResponse(holidys))
}

func CreateHoliday(c *gin.Context) {
	var input request.HolidayRequest

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusOK, gin.H{"errors": err})
		c.JSON(http.StatusOK, gin.H{"errors": errs})
		return
	}
	userIDValue, _ := c.Get("userID")
	parsedDate, _ := time.Parse("2006-01-02", input.Date)
	var holiday = models.Holiday{
		Name:      input.Name,
		Date:      parsedDate,
		CreatedBy: userIDValue.(uint),
	}
	if err := database.DB.Create(&holiday).Error; err != nil {
		c.JSON(http.StatusBadRequest, validation.DataBaseValidationError(err))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(holiday))
}

func UpdateHoliday(c *gin.Context) {
	id := c.Param("id")

	var input request.HolidayRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		c.JSON(http.StatusOK, gin.H{"errors": errs})
		return
	}
	var holiday models.Holiday
	if err := database.DB.First(&holiday, id).Error; err != nil {
		c.JSON(http.StatusOK, response.ErrorMessage("message", "Data not found"))
		return
	}
	parsedDate, _ := time.Parse("2006-01-02", input.Date)
	userIDValue, _ := c.Get("userID")
	holiday.Name = input.Name
	holiday.Date = parsedDate
	holiday.Status = input.Status
	holiday.CreatedBy = userIDValue.(uint)
	database.DB.Save(&holiday)
	c.JSON(http.StatusOK, response.SuccessResponse(&holiday))
}

func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")

	var holiday models.Holiday
	if err := database.DB.First(&holiday, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage("message", "Data not found"))
		return
	}
	database.DB.Delete(&holiday)
	c.JSON(http.StatusOK, response.SuccessResponse("Successfully delete"))
}
