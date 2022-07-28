package router

import (
	"fanama/auth/interface/routeHandler"

	"github.com/gofiber/fiber/v2"
)

func CreateAccessibleRouter(handler fiber.Router) {
	handler.Get("/", routeHandler.Accessible)
}
