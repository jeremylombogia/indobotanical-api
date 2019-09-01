package user

import (
	"log"
	"net/http"
	"time"

	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/internal"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"github.com/dgrijalva/jwt-go"
	"github.com/jeremylombogia/indobotanical-api/models"
	"github.com/labstack/echo"
)

// Login function
// It compare request email and password
// It compare between email and password to FindByEmailAndPassword() repositories
func Login(c echo.Context) error {
	payload := new(Payload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(401, err.Error())
	}

	var user, err = FindByEmail(payload)
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(401, internal.ErrorResponse{401, "We can't find your email in our database"})
		}

		return c.JSON(500, err.Error())
	}

	var requestPassword = []byte(payload.Data.Password)
	if !comparePasswords(user.Password, requestPassword) {
		return c.JSON(400, internal.ErrorResponse{400, "You enter a wrong password"})
	}

	var t, _ = generateToken(&user)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login success",
		"token":   t,
	})
}

// Register function
// It POST all user attribute in models
// It check wether duplicate email or password
// It response with user created response
func Register(c echo.Context) (err error) {
	payload := new(Payload)

	if err := c.Bind(payload); err != nil {
		return c.JSON(401, err.Error())
	}

	// Check same email found
	if _, err := FindByEmail(payload); err == nil {
		return c.JSON(409, internal.ErrorResponse{409, "Your email are already registered before"})
	}

	var user = models.User{}
	user = payload.Data.User
	user.ID = bson.NewObjectId()
	user.Level = 0 // 0 means it regular user
	user.CreatedAt = time.Now()

	// Hash Password
	var password = []byte(payload.Data.Password)
	user.Password = hashPassword(password)

	if user, err = StoreUser(&user); err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User created",
	})
}

// hashPassword
func hashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
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

// generateToken, it generate from echo token
func generateToken(user *models.User) (string, error) {
	// Create token
	var token = jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
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
