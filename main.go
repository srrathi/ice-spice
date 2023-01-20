package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/srrathi/ice-spice/controller"
	"github.com/srrathi/ice-spice/models"
	"github.com/srrathi/ice-spice/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load database")
	}

	err = models.MigrateDatabases(db)
	if err != nil {
		log.Fatal("Could not migrate database", err.Error())
	}

	r := controller.Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
