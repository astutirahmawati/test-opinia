package handlers

import "github.com/labstack/echo/v4"

type HandlerType interface {
	CreateType() echo.HandlerFunc
	GetbyID() echo.HandlerFunc
	UpdateType() echo.HandlerFunc
	DeleteType() echo.HandlerFunc
}
