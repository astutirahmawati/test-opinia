package handlers

import "github.com/labstack/echo/v4"

type HandlerPost interface {
	CreatePost() echo.HandlerFunc
	GetPostbyID() echo.HandlerFunc
	UpdatePost() echo.HandlerFunc
	DeletePost() echo.HandlerFunc
}
