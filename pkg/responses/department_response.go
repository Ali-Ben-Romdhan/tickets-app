package responses

import (
    "github.com/labstack/echo/v4"
)

type DepartmentResponse struct {
    Status  int       `json:"status"`
    Message string    `json:"message"`
    Data    *echo.Map `json:"data"`
}