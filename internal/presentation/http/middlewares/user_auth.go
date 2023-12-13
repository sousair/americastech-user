package http_middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	crypto_provider "github.com/sousair/americastech-user/internal/infra/cryptography"
)

func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Authorization header missing",
			})
		}

		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Authorization header malformed",
			})
		}

		jwtToken := authHeaderParts[1]

		cryptoProvider := crypto_provider.NewCryptoProvider()

		payload, err := cryptoProvider.VerifyAuthToken(jwtToken)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		c.Set("user_id", payload["id"].(string))

		return next(c)
	}
}
