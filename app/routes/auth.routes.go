package routes

import (
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/controllers"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	route.Post("/login", controllers.Login)
	route.Delete("/logout", controllers.Logout)
	route.Get("/me", middlewares.Protected(), controllers.Me)
	route.Post("/register", controllers.Register)
	route.Get("/refresh-token", controllers.RefreshToken)
}
