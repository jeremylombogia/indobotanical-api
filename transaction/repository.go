package transaction

import (
	"os"

	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/models"
	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "transactions"

var session, _ = config.MongoConnect()
var collection = session.DB("heroku_ghdkzt3f").C(DOCUMENT)

func FetchTransaction() ([]models.Transactions, error) {
	var transaction []models.Transactions

	var err = collection.Find(bson.M{}).All(&transaction)

	return transaction, err
}

func StoreTransation(transaction *models.Transactions) (models.Transactions, error) {
	var err = collection.Insert(*transaction)

	return *transaction, err
}
