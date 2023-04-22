package main

import (
	"log"
	"testgoapi/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	route.Register(app)

	log.Println(app.Listen(":8080"))
}
