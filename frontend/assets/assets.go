package assets

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Static("/assets", "frontend/assets")
}
