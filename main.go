package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/teerapoom/mini_api003/repository"
	"github.com/teerapoom/mini_api003/schema"
	"github.com/teerapoom/mini_api003/service"
)

func main() {
	londEnv()
	loadDatebase()
	process()

}

func process() {
	app := fiber.New()

	app.Post("/add/card", service.CreateCreditCard)
	app.Get("/view/user", service.GetUser)

	app.Listen(":8080")
}

func londEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatebase() {
	repository.InitDb()
	repository.Db.AutoMigrate(&schema.User{})
	repository.Db.AutoMigrate(&schema.CreditCard{})
}
