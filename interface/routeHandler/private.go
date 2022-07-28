package routeHandler

import "github.com/gofiber/fiber/v2"

func Restricted(c *fiber.Ctx) error {

	return c.SendString("Welcome to VIP")
}
