package main

import (
	"be_latihan/config"
	_ "be_latihan/docs"
	"be_latihan/model"
	"be_latihan/router"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title API be_latihan Praktikum 13
// @version 1.0
// @description Dokumentasi Swagger untuk backend be_latihan menggunakan Golang Fiber, GORM, PostgreSQL, dan JWT Bearer Token.
// @termsOfService http://swagger.io/terms/
// @contact.name Praktikum Pemrograman III
// @contact.email praktikum@example.com
// @host 127.0.0.1:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format: Bearer {token}
func main() {
	app := fiber.New()

	//swager
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "127.0.0.1:3000"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	app.Use(logger.New())

	config.InitDB()
	config.GetDB().AutoMigrate(&model.Mahasiswa{}, &model.User{})
	router.SetupRoutes(app)

	app.Listen(":3000")
}
