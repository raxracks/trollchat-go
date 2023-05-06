package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raxracks/trollchat/frontend/views"
	"github.com/raxracks/trollchat/frontend/assets"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("v1")
	views.RegisterRoutes(v1)
	assets.Setup(app)

	app.Listen(":8080")
}
