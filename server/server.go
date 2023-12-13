package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qahta0/image-processing-service/routes"
	"github.com/qahta0/image-processing-service/tasks"
	"log"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	app := fiber.New()
	routes.Setup(app)
	tasks.Init(redisAddress)
	defer tasks.Close()
	log.Fatal(app.Listen(":3000"))
}
