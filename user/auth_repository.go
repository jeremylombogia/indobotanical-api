package user

import (
	"os"

	"indobotanical-api/config"

	"indobotanical-api/models"

	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "users"

var sessionDb = config.MongoConnect()
var collection = sessionDb.DB(os.Getenv("DB_NAME")).C(DOCUMENT)

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
