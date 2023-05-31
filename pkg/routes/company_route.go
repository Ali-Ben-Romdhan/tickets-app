package routes

import (
    "echo-mongo-api/pkg/controllers" 

    "github.com/labstack/echo/v4"
)

func CompanyRoute(e *echo.Echo) {
    e.POST("/company", controllers.CreateCompany) 
		e.GET("/company/:companyId", controllers.GetACompany) 
		e.PUT("/company/:companyId", controllers.EditACompany) 
		e.DELETE("/company/:companyId", controllers.DeleteACompany) 
		e.GET("/companies", controllers.GetAllCompanies)
}

