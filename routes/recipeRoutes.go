package routes

import (
	"go_fun/controllers"
	"go_fun/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRecipeRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/recipes", middleware.AuthRequired(), controllers.CreateRecipe)
	api.Get("/recipes", controllers.GetRecipes)

}