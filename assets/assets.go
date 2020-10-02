package assets

import (
	"github.com/mdmims/go-fiber-api/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Assets struct for db model
type Assets struct {
	gorm.Model
	Name  string `json:"name"`
	Path string `json:"path"`
}
 
// GetAssets returns all assets
func GetAssets(c *fiber.Ctx) {
	db := database.DbConn
	var assets []Assets
	db.Find(&assets)
	c.JSON(assets)
}

// GetAsset returns a single asset
func GetAsset(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DbConn
	var asset Assets
	db.Find(&asset, id)
	c.JSON(asset)
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
