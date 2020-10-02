package main

import (
	"fmt"

	"github.com/mdmims/go-fiber-api/assets"
	"github.com/mdmims/go-fiber-api/database"
	
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DbConn, err = gorm.Open("sqlite3", "assets.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DbConn.AutoMigrate(&assets.Assets{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/assets", assets.GetAssets)	
	app.Get("/api/v1/asset/:id", assets.GetAsset)
	app.Post("/api/v1/asset", assets.NewAsset)
	app.Delete("/api/v1/asset/:id", assets.DeleteAsset)	
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DbConn.Close()
	setupRoutes(app)
	app.Listen(3000)
}
