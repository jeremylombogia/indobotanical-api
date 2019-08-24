package transaction

import (
	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/models"
)

const DOCUMENT string = "transactions"

var session, _ = config.MongoConnect()
var collection = session.DB(config.COLLECTION).C(DOCUMENT)

func StoreTransation(transaction *models.Transactions) (models.Transactions, error) {
	var err = collection.Insert(*transaction)

	return *transaction, err
}
