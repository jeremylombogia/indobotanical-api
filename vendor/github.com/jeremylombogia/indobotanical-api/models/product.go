package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Products struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Stock       int           `json:"stock" bson:"stock"`
	Price       int           `json:"price" bson:"price"`
	Thumbnail   string        `json:"thumbnail" bson:"thumbnail,omitempty" `
	Avaibility  bool          `json:"avaibility" bson:"avaibility"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt" bson:"updatedAt"`
}
