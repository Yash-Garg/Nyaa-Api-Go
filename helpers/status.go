package helpers

import "github.com/gofiber/fiber/v2"

func Status(ctx *fiber.Ctx) error {
	return ctx.SendString("Nyaa API v1 // Alive")
}
