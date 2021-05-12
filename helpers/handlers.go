package helpers

import (
	"github.com/Yash-Garg/nyaa-api-go/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AliveHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Nyaa API v1 // Alive")
}

func GroupHandler(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func CheckNyaaUrl() string {
	_, err := http.Get(utils.NyaaBaseUrl)
	if err != nil {
		return utils.NyaaAltUrl
	} else {
		return utils.NyaaBaseUrl
	}
}
