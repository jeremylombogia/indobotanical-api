package user

import (
	"os"

	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/models"
	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "users"

var session, _ = config.MongoConnect()
var collection = session.DB(os.Getenv("COLLECTION_NAME")).C(DOCUMENT)

// FindByEmail to see duplicate in register
// TODO:: Refactor this to one function same as FindByEmailAndPassword()
func FindByEmail(payload *Payload) (models.User, error) {
	var user models.User

	var err = collection.Find(bson.M{
		"email": payload.Data.Email,
	}).One(&user)

	return user, err
}

func StoreUser(user *models.User) (models.User, error) {
	var err = collection.Insert(*user)

	return *user, err
}
