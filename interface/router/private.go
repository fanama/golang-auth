package router

import (
	"fanama/auth/interface/routeHandler"

	"github.com/gofiber/fiber/v2"
)

func CreatePrivateRouter(handler fiber.Router) {
	handler.Get("/", routeHandler.Restricted)
}
