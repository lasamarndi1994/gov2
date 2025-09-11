package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lasamarndi1994/gov2/api/handler"
	"github.com/lasamarndi1994/gov2/api/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	api.POST("/login", handler.HandleLogin)
	api.POST("/register", handler.HandleRegister)
	api.POST("/forgot-password", handler.HandleForgotPassword)
	api.POST("/reset-password", handler.HandleResetPassword)

	//auth := api.Group("/", middleware.AuthMiddleware)

	api.Use(middleware.AuthMiddleware)
	{
		api.GET("/employees", handler.GetUserDetails)
		api.POST("/create-employee", handler.CreateEmployee)
		//Department
		api.GET("/departments", handler.GetDepartments)
		api.POST("/department", handler.CreateDepartment)
		api.GET("department/:id", handler.ShowDepartment)
		api.PUT("/department/:id", handler.UpdateDepartment)
		api.DELETE("department/:id", handler.DeleteDepartment)

		//Designation
		api.GET("/designations", handler.GetDesignations)
		api.POST("/designation", handler.CreateDesignation)
		api.GET("designation/:id", handler.ShowDesignation)
		api.PUT("/designation/:id", handler.UpdateDesignation)
		api.DELETE("designation/:id", handler.DeleteDesignation)

		//attendance
		api.GET("attendences", handler.GetAttendance)
		api.POST("/attendance", handler.StoreAttendance)
		api.PUT("/attendance", handler.UpdateAttendance)

		//Leave
		api.POST("/create-leave-type", handler.CreateLeave)
		api.GET("/get-leave-details", handler.GetLeaveDetails)
		api.POST("/apply-leave", handler.ApplyLeave)
		api.PUT("/update-leave/:id", handler.UpdateLeave)
		api.DELETE("delete-leave/:id", handler.DeleteLeave)

		//holiday
		api.GET("/get-holidays", handler.GetHolidays)
		api.POST("/create-holiday", handler.CreateHoliday)
		api.PUT("/update-holiday/:id", handler.UpdateHoliday)
		api.DELETE("/delete-holiday/:id", handler.DeleteHoliday)

		//salary
		api.GET("/get-emp-salary", handler.GetSalary)
		api.POST("/create-emp-slarly", handler.CreateSalary)
		api.PUT("/update-emp-slarly/:id", handler.UpdateSalary)
		api.DELETE("/delete-emp-slarly/:id", handler.DeleteSalary)

	}
	return router
}
