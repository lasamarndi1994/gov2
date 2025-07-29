package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/models"
	"github.com/lasamarndi1994/gov2/internal/database"
	"github.com/lasamarndi1994/gov2/utility/response"
)

func GetUserDetails(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, response.ResponseData("Success", users))
}

func StoreUser(c *gin.Context) {
	var user models.User
	database.DB.Create(&user)

}
