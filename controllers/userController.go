package controllers

import (
	"context"
	"go_fun/config"
	"go_fun/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"bad request"})
	}

	collection := config.DB.Collection("users")

	var user models.User

	err = collection.FindOne(context.Background(), bson.M{"_id":objectId}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error":"user not found"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func GetUserRecipes(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"invalid user ID"})
	}

	collection := config.DB.Collection("recipes")

	var recipes []models.Recipe
	cursor, err := collection.Find(context.Background(), bson.M{"author_id":objectId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"internal server error"})
	}

	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &recipes); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"internal server error"})
	}

	return c.Status(http.StatusOK).JSON(recipes)

}