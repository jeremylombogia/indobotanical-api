package user

import (
	"indobotanical-api/internal"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

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
	if !internal.ComparePassword(user.Password, requestPassword) {
		return c.JSON(400, internal.ErrorResponse{400, "You enter a wrong password"})
	}

	var t, _ = internal.GenerateToken(&user)

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
	user.Password = internal.HashPassword(password)

	if user, err = StoreUser(&user); err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User created",
	})
}
