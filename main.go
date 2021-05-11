package main

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"os"

	"github.com/Yash-Garg/nyaa-api-go/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
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

	app.Get("/", helpers.AliveHandler)

	v1 := app.Group("/nyaa", helpers.GroupHandler)

	v1.Get("/:category/:sub_category?", helpers.GetNyaa)
	v1.Get("/id/:id", helpers.GetUploadInfo)
	v1.Get("/user/:userID", helpers.GetUserUploads)
	_ = app.Listen(getPort())
}
