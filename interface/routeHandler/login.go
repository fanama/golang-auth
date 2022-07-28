package routeHandler

import (
	"fanama/auth/domain"
	"fanama/auth/infra/security"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("user")
	password := c.FormValue("pass")

	user := domain.User{Name: username, Password: password}

	t, tr, err := security.GenerateToken(user)

	if err != nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	}

	accessCookie, refreshCookie := security.GetAuthCookies(t, tr)

	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.JSON(fiber.Map{"accesstoken": t, "refreshToken": tr})
}
func Logout(c *fiber.Ctx) error {

	access_cookie, refresh_cookie := security.InvalidToken(c.Cookies("access_token"), c.Cookies("refresh_token"))
	c.Cookie(access_cookie)
	c.Cookie(refresh_cookie)
	return c.Status(200).SendString("success")
}
