package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go_fun/config"
	"go_fun/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)
var collection *mongo.Collection
func main() {
    fmt.Println("Starting the application...")

    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file", err)
    }

    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowCredentials: true,
	}))

    config.ConnectDB()
    defer config.DisconnectDB()

    routes.SetupRoutes(app)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("\nGracefully shutting down...")
        config.DisconnectDB()
        os.Exit(0)
    }()

    port := os.Getenv("PORT")
    if port == "" {
        port = "4000"
    }

    log.Fatal(app.Listen("0.0.0.0:" + port))
}
