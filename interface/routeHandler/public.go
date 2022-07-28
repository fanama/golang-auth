package routeHandler

import "github.com/gofiber/fiber/v2"

func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}
