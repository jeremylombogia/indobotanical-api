package product

import (
	"os"
	"time"

	"github.com/jeremylombogia/indobotanical-api/config"
	"github.com/jeremylombogia/indobotanical-api/models"
	"gopkg.in/mgo.v2/bson"
)

const DOCUMENT string = "products"

var session, _ = config.MongoConnect()
var collection = session.DB(os.Getenv("DB_NAME")).C(DOCUMENT)

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

func UpdateProduct(product *models.Products, id string) (models.Products, error) {
	previousData, _ := FindProduct(id)

	if product.Name == "" {
		product.Name = previousData.Name
	}

	if product.Description == "" {
		product.Description = previousData.Description
	}

	if product.Thumbnail == "" {
		product.Thumbnail = previousData.Thumbnail
	}

	var payload = bson.M{
		"name":        product.Name,
		"description": product.Description,
		"stock":       product.Stock,
		"price":       product.Price,
		"thumbnail":   product.Thumbnail,
		"createdAt":   product.CreatedAt,
		"updatedAt":   time.Now(),
	}

	var err = collection.Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": payload},
	)

	return *product, err
}
