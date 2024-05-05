package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/srrathi/ice-spice/controller"
	"github.com/srrathi/ice-spice/models"
	"github.com/srrathi/ice-spice/storage"
)

func main() {
	APP_PORT := os.Getenv("PORT")
	if APP_PORT == "" {
		APP_PORT = "10000"
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

	app.Static("/", "./public")
	app.Static("/public", "./public")

	app.Use(limiter.New(limiter.Config{
		Max:            20,
		Expiration:     60 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	r.SetupRoutes(app)
	app.Listen(":" + APP_PORT)
}
