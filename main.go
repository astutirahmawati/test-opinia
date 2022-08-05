package main

import (
	"opinia/configs"
	"opinia/delivery/handlers"
	typeHandlers "opinia/delivery/handlers/post_type"
	postHandlers "opinia/delivery/handlers/postingan"
	userHandlers "opinia/delivery/handlers/user"
	"opinia/delivery/middlewares"
	"opinia/delivery/routes"
	"opinia/delivery/validations"
	typeRepository "opinia/repository/post_type"
	postRepository "opinia/repository/postingan"
	userRepository "opinia/repository/user"
	authService "opinia/services/authentification"
	typeService "opinia/services/post_type"
	postService "opinia/services/postingan"
	userService "opinia/services/user"
	utils "opinia/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database
	config := configs.Get()

	// Init DB
	DB := utils.NewMysqlGorm(config)

	// Migrate
	utils.Migrate(DB)

	// Initiate Echo
	e := echo.New()

	// repository
	userRepository := userRepository.NewrUserRepo(DB)
	typeRepository := typeRepository.NewTypeRepo(DB)
	postRepository := postRepository.NewPostRepo(DB)

	// Validation
	validation := validations.NewValidation(validator.New())
	// services
	authService := authService.NewAuthService(userRepository)

	userService := userService.NewUserService(userRepository)
	typeService := typeService.NewTypeService(typeRepository)
	postService := postService.NewPostService(postRepository)
	// delivery
	authHandler := handlers.NewAuthHandler(authService, validation)
	userHandler := userHandlers.NewUserHandler(userService, validation)
	typeHandler := typeHandlers.NewTypeHandler(typeService)
	postHandler := postHandlers.NewPostHandler(postService, validation)
	// Routes
	routes.AuthRoute(e, authHandler)
	routes.UserRoute(e, *userHandler)
	routes.TypeRoute(e, *typeHandler)
	routes.PostRoute(e, *postHandler)

	// Middlewares
	middlewares.General(e)

	e.Logger.Fatal(e.Start(":8000"))
}
