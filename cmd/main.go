package main

import (
	"echo-mongo-api/pkg/configs"
	"echo-mongo-api/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    //run database
    configs.ConnectDB()
    //routes
    routes.UserRoute(e) 
    routes.DepartmentRoute(e) 
    routes.CompanyRoute(e) 

    e.Logger.Fatal(e.Start(":6000"))
}