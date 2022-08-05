package handlers

import (
	"opinia/delivery/helpers"
	"opinia/delivery/middlewares"
	"strconv"

	"net/http"
	"opinia/entities"
	typeService "opinia/services/post_type"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TypeHandler struct {
	typeService typeService.TypeServiceInterface
}

func NewTypeHandler(service typeService.TypeServiceInterface) *TypeHandler {
	return &TypeHandler{
		typeService: service,
	}
}

func (th *TypeHandler) CreateType() echo.HandlerFunc {
	return func(c echo.Context) error {

		var req entities.PostType
		err := c.Bind(&req)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		userId, role := middlewares.ExtractTokenRoleID(c)

		if err != nil || role != "admin" {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		add, err := th.typeService.CreateType(uint(userId), req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		// response
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+add.Jenis, add))
	}
}

func (th *TypeHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idType, _ := strconv.Atoi(c.Param("id"))

		res, err := th.typeService.GetbyID(uint(idType))
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Not Found"))
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data "+res.Jenis, res))
	}
}

func (th *TypeHandler) UpdateType() echo.HandlerFunc {
	return func(c echo.Context) error {

		updateReq := entities.PostType{}
		c.Bind(&updateReq)

		id, err := strconv.Atoi(c.Param("id"))
		_, role := middlewares.ExtractTokenRoleID(c)

		if role != "admin" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		userRes, err := th.typeService.UpdateType(uint(id), updateReq)
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Not Found"))
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update Customer", userRes))
	}
}

func (th *TypeHandler) DeleteType() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		userId, role := middlewares.ExtractTokenRoleID(c)
		if role != "admin" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}
		err = th.typeService.DeleteType(uint(userId), uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Not Found"))
		}
		// response
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}
