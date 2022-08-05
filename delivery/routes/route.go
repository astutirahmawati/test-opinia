package routes

import (
	"opinia/delivery/handlers"
	typee "opinia/delivery/handlers/post_type"
	post "opinia/delivery/handlers/postingan"
	user "opinia/delivery/handlers/user"
	"opinia/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, u user.UserHandler) {

	customerGroup := e.Group("/customer")

	customerGroup.POST("", u.CreateCustomer())
	customerGroup.GET("/:id", u.GetByID())
	customerGroup.PUT("/:id", u.UpdateCustomer(), middlewares.JWTMiddleware())
	customerGroup.DELETE("/:id", u.DeleteCustomer(), middlewares.JWTMiddleware())
}

func TypeRoute(e *echo.Echo, t typee.TypeHandler) {

	typeGroup := e.Group("/type")

	typeGroup.POST("", t.CreateType(), middlewares.JWTMiddleware())
	typeGroup.GET("/:id", t.GetByID())
	typeGroup.PUT("/:id", t.UpdateType(), middlewares.JWTMiddleware())
	typeGroup.DELETE("/:id", t.DeleteType(), middlewares.JWTMiddleware())
}
func PostRoute(e *echo.Echo, p post.PostHandler) {

	typeGroup := e.Group("/post")

	typeGroup.POST("", p.CreatePost(), middlewares.JWTMiddleware())
	typeGroup.GET("/:id", p.GetPostByID())
	typeGroup.PUT("/:id", p.UpdatePost(), middlewares.JWTMiddleware())
	typeGroup.DELETE("/:id", p.DeletePost(), middlewares.JWTMiddleware())
}
func AuthRoute(e *echo.Echo, l *handlers.AuthHandler) {
	e.POST("/login", l.Login())

}
