package handlers

import (
	"net/http"
	"opinia/delivery/helpers"
	validation "opinia/delivery/validations"
	"opinia/entities"
	authService "opinia/services/authentification"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AuthHandler struct {
	authService *authService.AuthService
	valid       validation.Validation
}

func NewAuthHandler(service *authService.AuthService, Valid validation.Validation) *AuthHandler {
	return &AuthHandler{
		authService: service,
		valid:       Valid,
	}
}

func (handler AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Populate request input
		var userReq entities.AuthRequest
		err := c.Bind(&userReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = handler.valid.Validation(userReq)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		// define link hateoas
		// call auth service login
		authRes, err := handler.authService.Login(userReq)
		if err != nil {
			if err.Error() == "invalid username/password" {
				return c.JSON(http.StatusForbidden, helpers.ErrorAuthorize())
			}
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("User Not Found"))
		}

		// send response
		return c.JSON(http.StatusOK, helpers.LoginOK(authRes))
	}
}
