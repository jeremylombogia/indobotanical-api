package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transactions struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Products
	User
	TotalPrice int       `json:"totalPrice" bson:"totalPrice"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}
