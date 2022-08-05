package handlers

import "github.com/labstack/echo/v4"

type HandleUser interface {
	// CreateInternal(c echo.Context) error
	CreateCustomer() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAllMember() echo.HandlerFunc
	UpdateCustomer() echo.HandlerFunc
	DeleteCustomer() echo.HandlerFunc
}
