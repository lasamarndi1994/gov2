package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/internal/request"
	"github.com/lasamarndi1994/gov2/utility/validation"
)

func GetSalary(c *gin.Context) {

}

func CreateSalary(c *gin.Context) {
	var input request.SalaryRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		errs := validation.FormatValidationError(err)
		if errs != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"erros": errs})
		}
	}
	userIDValue, _ := c.Get("userID")
	var salary = models.SalaryPay{
		UserId:              userIDValue.(uint),
		BasicSalary:         input.BasicSalary,
		HouseRentAllowance:  input.HouseRentAllowance,
		ConveyanceAllowance: input.ConveyanceAllowance,
		MedicalAllowance:    input.MedicalAllowance,
		DearnessAllowance:   input.DearnessAllowance,
		Allowance:           input.Allowance,
		OtherAllowances:     input.OtherAllowances,
	}
	var salary_deducation = models.SalaryDeduction{
		UserId:          userIDValue.(uint),
		ProfessionalTax: input.ProfessionalTax,
		IncomeTax:       input.IncomeTax,
		Leave:           input.Leave,
		ProvidentFund:   input.ProvidentFund,
		OtherDeduction:  input.OtherDeduction,
	}
	database.DB.Create(&salary)
	database.DB.Create(&salary_deducation)
}

func UpdateSalary(c *gin.Context) {
}

func DeleteSalary(c *gin.Context) {
}
