package routes

import (
    "echo-mongo-api/pkg/controllers" 

    "github.com/labstack/echo/v4"
)

func DepartmentRoute(e *echo.Echo) {
    e.POST("/department", controllers.CreateDepartment) 
		e.GET("/department/:userId", controllers.GetADepartment) 
		e.PUT("/department/:userId", controllers.EditADepartment) 
		e.DELETE("/department/:userId", controllers.DeleteADepartment) 
		e.GET("/departments", controllers.GetAllDepartments)
}

