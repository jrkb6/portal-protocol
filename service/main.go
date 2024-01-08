package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"service/config"
	IdentityService "service/controllers"
	"service/routers"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./build")
	config.LoadConfig()
	IdentityService.InitOwnerKeys()
	routers.Api(app)

	log.Fatal(app.Listen(":3000"))

}
