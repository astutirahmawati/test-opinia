package handlers

import (
	"opinia/delivery/helpers"
	"opinia/delivery/middlewares"
	"strconv"

	"net/http"
	validation "opinia/delivery/validations"
	"opinia/entities"
	postService "opinia/services/postingan"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PostHandler struct {
	postService postService.PostServiceInterface
	valid       validation.Validation
}

func NewPostHandler(service postService.PostServiceInterface, Valid validation.Validation) *PostHandler {
	return &PostHandler{
		postService: service,
		valid:       Valid,
	}
}

func (ph *PostHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {

		var req entities.Postingan
		err := c.Bind(&req)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequestBind(err))
		}

		err = ph.valid.Validation(req)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}
		userId, role := middlewares.ExtractTokenRoleID(c)

		if err != nil || role != "customer" {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		add, err := ph.postService.CreatePost(uint(userId), req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.StatusBadRequest(err))
		}

		// response
		return c.JSON(http.StatusCreated, helpers.StatusCreate("Success Create "+add.Title, add))
	}
}

func (ph *PostHandler) GetPostByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idPost, _ := strconv.Atoi(c.Param("id"))

		res, err := ph.postService.GetbyID(uint(idPost))
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Not Found"))
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusGetDataID("Success Get Data"+res.Title, res))
	}
}

func (ph *PostHandler) UpdatePost() echo.HandlerFunc {
	return func(c echo.Context) error {

		updateReq := entities.Postingan{}
		c.Bind(&updateReq)

		id, err := strconv.Atoi(c.Param("id"))
		userId, role := middlewares.ExtractTokenRoleID(c)

		if role != "customer" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}

		userRes, err := ph.postService.UpdatePost(uint(userId), uint(id), updateReq)
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound("Not Found"))
		}

		// response
		return c.JSON(http.StatusOK, helpers.StatusUpdate("Success Update ", userRes))
	}
}

func (ph *PostHandler) DeletePost() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		userId, role := middlewares.ExtractTokenRoleID(c)
		if role != "customer" || err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ErrorAuthorize())
		}
		err = ph.postService.DeletePost(uint(userId), uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, helpers.StatusNotFound(" Not Found"))
		}
		// response
		return c.JSON(http.StatusOK, helpers.StatusDelete())
	}
}
