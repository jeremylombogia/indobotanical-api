package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transactions struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Products     bson.ObjectId `bson:"productId" json:"productId"`
	User         bson.ObjectId `bson:"userId" json:"userId"`
	TotalPrice   int           `json:"totalPrice" bson:"totalPrice"`
	Status       int           `json:"status" bson:"status"` // 0 means unpaid, 1 payment proof uploaded, 2 means paid
	PaymentProof interface{}   `json:"paymentProof" bson:"paymentProof"`
	CreatedAt    time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt" bson:"updatedAt"`
}
