package main

import (
	"fanama/auth/interface/routeHandler"
	"fanama/auth/interface/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Login route
	router.CreateLoginRouter(app.Group("/login"))
	router.CreateLogoutRouter(app.Group("/logout"))

	// Unauthenticated route
	router.CreateAccessibleRouter(app.Group("/public"))
	app.Use(routeHandler.CheckToken)

	// Restricted Routes

	router.CreatePrivateRouter(app.Group("/api"))

	app.Listen(":3000")
}
