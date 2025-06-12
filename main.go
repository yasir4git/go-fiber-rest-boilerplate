package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/IKHINtech/go-fiber-rest-boilerplate/config"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/database"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/middlewares"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/routes"
)

// @title Go Fiber Boilerplate
// @version 1.0
// @description This is a sample boilerplate for Go Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5001
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.LoadConfig()
	database.ConnectDB()

	app := fiber.New(
		fiber.Config{
			AppName: "Go Fiber Boilerplate",
		},
	)

	middlewares.SetupRecovery(app)
	middlewares.SetupCORS(app)

	routes.SetupRoutesApp(app)

	if err := app.Listen(":" + config.AppConfig.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
