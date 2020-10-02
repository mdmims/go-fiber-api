package main

import (
	"github.com/mdmims/go-fiber-api/assets"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/assets", assets.GetAssets)	
	app.Get("/api/v1/asset/:id", assets.GetAsset)
	app.Post("/api/v1/asset", assets.NewAsset)
	app.Delete("/api/v1/asset/:id", assets.DeleteAsset)	
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}