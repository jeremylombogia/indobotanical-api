package user

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jeremylombogia/indobotanical-api/models"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	payload := new(Payload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(401, err.Error())
	}

	var user, err = FindByEmailAndPassword(payload.Data.Email, payload.Data.Password)
	if err != nil {
		if err.Error() == "not found" {
			return echo.ErrUnauthorized
		}

		return c.JSON(500, err.Error())
	}

	var t, _ = generateToken(&user)

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func generateToken(user *models.User) (string, error) {
	// Create token
	var token = jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["level"] = user.Level
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	// TODO:: change "secret" to random code that write in ENV
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, err
}
