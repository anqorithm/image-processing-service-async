package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qahta0/image-processing-service/handlers"
)

func Setup(app *fiber.App) {
	app.Post("/process-image", handlers.UploadImage)
}
