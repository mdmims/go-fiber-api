package assets

import (
	"github.com/gofiber/fiber"	
)

// GetAssets returns all assets
func GetAssets(c *fiber.Ctx) {
	c.Send("all assets")
}

// GetAsset returns a single asset
func GetAsset(c *fiber.Ctx) {
	c.Send("one asset")
}

// DeleteAsset deletes a single asset
func DeleteAsset(c *fiber.Ctx) {
	c.Send("delete asset")
}

// NewAsset creates a new asset
func NewAsset(c *fiber.Ctx) {
	c.Send("create asset")
}

// TODO add update endpoint for assets
