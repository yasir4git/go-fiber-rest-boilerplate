package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/controllers"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/dto"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/repositories"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/services"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/middlewares"
)

func UserRoutes(route fiber.Router) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	route.Get("/", userController.GetUsers)
	route.Get("/:id", userController.GetUser)
	route.Post("/", middlewares.ValidateRequest(&dto.CreateUserDTO{}), userController.CreateUser)
	route.Patch(
		"/:id",
		middlewares.ValidateRequest(&dto.UpdateUserDTO{}),
		userController.UpdateUser,
	)
	route.Delete("/:id", userController.DeleteUser)
}
