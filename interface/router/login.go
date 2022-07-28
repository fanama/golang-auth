package router

import (
	"fanama/auth/interface/routeHandler"

	"github.com/gofiber/fiber/v2"
)

func CreateLoginRouter(handler fiber.Router) {
	handler.Post("/", routeHandler.Login)
}
