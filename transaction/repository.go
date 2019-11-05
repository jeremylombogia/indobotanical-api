package transaction

import (
	"os"

	"indobotanical-api/config"
	"indobotanical-api/models"

	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "transactions"

var sessionDb = config.MongoConnect()
var collection = sessionDb.DB(os.Getenv("DB_NAME")).C(DOCUMENT)

//func FetchTransaction() ([]models.Transactions, error) {
//	var transaction []models.Transactions
//
//	var err = collection.Find(bson.M{}).All(&transaction)
//
//	return transaction, err
//}

func FetchTransactionsByUserID(userID string) ([]models.Transactions, error) {
	var transaction []models.Transactions

	var err = collection.Find(bson.M{
		"userId": userID,
	}).All(&transaction)

	return transaction, err
}

func StoreTransation(transaction *models.Transactions) (models.Transactions, error) {
	var err = collection.Insert(*transaction)

	return *transaction, err
}
