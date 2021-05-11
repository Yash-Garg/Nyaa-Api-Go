package helpers

import "github.com/gofiber/fiber/v2"

func AliveStatus(ctx *fiber.Ctx) error {
	return ctx.SendString("Nyaa API v1 // Alive")
}
