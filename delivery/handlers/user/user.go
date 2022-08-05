package handlers

import (
	"fmt"

	"opinia/delivery/helpers"
	"strconv"

	"net/http"
	middleware "opinia/delivery/middlewares"
	validation "opinia/delivery/validations"
	"opinia/entities"
	userService "opinia/services/user"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	userService userService.UserServiceInterface
	valid       validation.Validation
}

func NewUserHandler(service userService.UserServiceInterface, Valid validation.Validation) *UserHandler {
	return &UserHandler{
		userService: service,
		valid:       Valid,
	}
}

func (handler *UserHandler) CreateInternal() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request ke user request
		var userReq entities.CreateUserRequest
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
		token := c.Get("user")
		_, role, err := middleware.ReadToken(token)

		if err != nil || role != "admin" {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		hashedPassword, _ := helpers.HashPassword(userReq.Password)
		userReq.Password = hashedPassword
		// registrasi user via call user service
		user, err := handler.userService.CreateUser(userReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ErrorRegister("Email Be Registered"))
		}
		// Konversi hasil repository menjadi user response
		userRes := entities.InternalResponse{}
		copier.Copy(&userRes, &user)

		// generate token
		Token, err := middleware.CreateToken(int(user.ID), user.Name, user.Role)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// Buat auth response untuk dimasukkan token dan user
		authRes := entities.InternalAuthResponse{
			Token: Token,
			User:  userRes,
		}

		// response
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+user.Role, authRes))
	}
}

func (handler *UserHandler) CreateCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request ke user request
		var userReq entities.CreateUserRequest
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

		hashedPassword, _ := helpers.HashPassword(userReq.Password)
		userReq.Password = hashedPassword

		// registrasi user via call user service
		user, err := handler.userService.CreateUser(userReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ErrorRegister("Email Be Registered"))
		}
		// Konversi hasil repository menjadi user response
		userRes := entities.InternalResponse{}
		copier.Copy(&userRes, &user)
		fmt.Println(user)

		// generate token
		token, err := middleware.CreateToken(int(user.ID), user.Name, user.Role)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, helpers.InternalServerError())
		}

		// Buat auth response untuk dimasukkan token dan user
		authRes := entities.InternalAuthResponse{
			Token: token,
			User:  userRes,
		}
		// response
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create ", authRes))
	}
}

func (handler *UserHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idUser, _ := strconv.Atoi(c.Param("id"))
		// Update via user service call

		userRes, err := handler.userService.GetbyID(uint(idUser))
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("User By ID Not Found"))
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data"+userRes.Name, userRes))
	}
}

func (handler *UserHandler) UpdateCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request to user request

		userReq := entities.UpdateCustomerRequest{}
		c.Bind(&userReq)

		// Get token
		token := c.Get("user")
		id, _ := strconv.Atoi(c.Param("id"))
		idToken, role, err := middleware.ReadToken(token)

		if id != idToken || role != "customer" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		// Update via user service call
		userRes, err := handler.userService.UpdateCustomer(uint(id), userReq)
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("User By ID Not Found"))
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Customer", userRes))
	}
}

func (handler *UserHandler) DeleteCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Get("user")
		id, _ := strconv.Atoi(c.Param("id"))
		idToken, role, err := middleware.ReadToken(token)
		if id != idToken || role != "customer" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}
		err = handler.userService.DeleteCustomer(uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("User By ID Not Found"))
		}
		// response
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}
