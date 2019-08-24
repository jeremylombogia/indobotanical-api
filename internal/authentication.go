package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetAuthenticatedUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	return id
}
