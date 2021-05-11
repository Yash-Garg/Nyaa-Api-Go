package helpers

import "github.com/gofiber/fiber/v2"

func AliveHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Nyaa API v1 // Alive")
}

func GroupHandler(ctx *fiber.Ctx) error {
	return ctx.Next()
}
