package main

import (
	"os"

	"github.com/Yash-Garg/nyaa-api-go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":6969"
}

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	app := fiber.New()

	app.Use(cors.New())
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "pass#1234",
		},
	}))

	app.Get("/", helpers.Status)
	app.Get("/nyaa", helpers.GetNyaa)
	_ = app.Listen(getPort())
}
