package internal

import (
	"log"
	"time"

	"indobotanical-api/config"
	"indobotanical-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetAuthenticatedUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	return id
}

// HashPassword, it hash password to bcrypt from plain string (byte in here)
func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePassword it compare password between plain string and hashed password in byte
func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GenerateToken, it generate from echo token
func GenerateToken(user *models.User) (string, error) {
	// Create token
	var token = jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["city"] = user.City
	claims["country"] = user.Country
	claims["address"] = user.Address
	claims["level"] = user.Level
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	// TODO:: change "secret" to random code that write in ENV
	t, err := token.SignedString([]byte(config.APPKEY))
	if err != nil {
		return "", err
	}

	return t, err
}
