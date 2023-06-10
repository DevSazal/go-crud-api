package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/DevSazal/go-crud-api/app/handlers"
	"github.com/DevSazal/go-crud-api/app/repositories"
	"github.com/DevSazal/go-crud-api/app/services"
)

func main() {
	app := fiber.New()

	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI(envMongo())
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	userRepository := repositories.NewUserRepository(client)

	// Initialize services
	userService := services.NewUserService(*userRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(*userService)

	// Routes
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:id", userHandler.GetUser)
	app.Post("/users", userHandler.CreateUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

func envMongo() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGO_URI")
}
