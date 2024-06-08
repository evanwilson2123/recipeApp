package routes

import (
	"go_fun/controllers"
	"go_fun/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/user/profile", middleware.AuthRequired(), controllers.GetUserProfile)
	api.Get("/user/recipes", middleware.AuthRequired(), controllers.GetUserRecipes)
}