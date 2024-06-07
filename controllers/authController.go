package controllers

import (
	"context"
	"go_fun/config"
	"go_fun/models"
	"go_fun/utils"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	collection := config.DB.Collection("users")

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"Invalid input"})
	}

	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		log.Printf("Error finding user: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"user already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"})
	}

	user.Password = string(hashedPassword)

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"failed to create user"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message":"User created successfully"})
}

func Login(c *fiber.Ctx) error {
	collection := config.DB.Collection("users")

	var input models.User
	if err := c.BodyParser(&input); err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"Invalid input"})
	}

	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	}

	jwtToken, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to generate token"})
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to generate refresh token"})
	}

	c.Cookie(&fiber.Cookie{
		Name: "refresh_token",
		Value: refreshToken,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{"token":jwtToken})
}

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Missing refresh token"})
	}

	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid refresh token"})
	}

	jwtToken, err := utils.GenerateJWT(claims.UserId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to generate token"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token":jwtToken})
}