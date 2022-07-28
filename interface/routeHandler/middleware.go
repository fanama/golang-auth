package routeHandler

import (
	"fanama/auth/infra/security"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckToken(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	refreshToken := c.Cookies("refresh_token")
	var err error

	token := c.Get("Authorization")
	tokenArray := strings.Split(token, " ")

	if token != "" && security.VerifyToken(tokenArray[len(tokenArray)-1]) == nil {
		return c.Next()
	}

	if refreshToken == "" {
		return c.Status(fiber.StatusForbidden).SendString("no refresh token found")
	}

	if accessToken == "" {
		accessToken, refreshToken, err = security.RenewToken(refreshToken)

		if err != nil {

			return c.Status(fiber.StatusForbidden).SendString("can't get new tokens")
		}
		accessTokenCookie, refreshTokenCoockie := security.GetAuthCookies(accessToken, refreshToken)
		c.Cookie(accessTokenCookie)
		c.Cookie(refreshTokenCoockie)

	}

	if security.VerifyToken(accessToken) != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("accessToken is not valid")
	}

	return c.Next()
}
