package product

import (
	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/models"
	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "products"

var session, _ = config.MongoConnect()
var collection = session.DB(config.COLLECTION).C(DOCUMENT)

func FetchProduct() ([]models.Products, error) {
	var product []models.Products

	var err = collection.Find(bson.M{}).All(&product)

	return product, err
}

func FindProduct(id string) (models.Products, error) {
	var product models.Products

	var err = collection.FindId(bson.ObjectIdHex(id)).One(&product)

	return product, err
}

func StoreProduct(product *models.Products) (models.Products, error) {
	var err = collection.Insert(*product)

	return *product, err
}

func UpdateProduct(id string, product *models.Products) (models.Products, error) {
	var err = collection.Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{
			"name":        "new Name",
			"description": product.Description,
		}},
	)

	return *product, err
}
