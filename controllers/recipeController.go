package controllers

import (
	"context"
	"go_fun/config"
	"log"
	"net/http"
	"time"

	"go_fun/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRecipe(c *fiber.Ctx) error {
    collection := config.DB.Collection("recipes")

    var recipe models.Recipe

    if err := c.BodyParser(&recipe); err != nil {
        log.Printf("Error parsing body: %v\n", err)
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

	userId := c.Locals("user_id").(string)
	authorId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Printf("Error converting userID to ObjectId: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"internal server error"})
	}

    recipe.ID = primitive.NewObjectID()
	recipe.AuthorID = authorId
    recipe.CreatedAt = time.Now().Unix()

    _, err = collection.InsertOne(context.Background(), recipe)
    if err != nil {
        log.Printf("Error inserting recipe: %v\n", err)
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create recipe"})
    }

    return c.Status(http.StatusCreated).JSON(recipe)
}

func GetRecipes(c *fiber.Ctx) error {
	collection := config.DB.Collection("recipes")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"internal server error"})
	}
	defer cursor.Close(context.Background())

	var recipes []models.Recipe
	if err := cursor.All(context.Background(), &recipes); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"internal server error"})
	}

	return c.Status(http.StatusOK).JSON(recipes)
}