package user

import (
	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/models"
	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "users"

var session, _ = config.MongoConnect()
var collection = session.DB(config.COLLECTION).C(DOCUMENT)

func FindByEmailAndPassword(email string, password string) (models.User, error) {
	var user models.User

	var err = collection.Find(bson.M{
		"email":    email,
		"password": password,
	}).One(&user)

	return user, err
}
