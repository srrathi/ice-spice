package main

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Repository struct{ 
	DB *gorm.DB
}

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	app := fiber.New()

}