package config

import (
	"os"

	"gopkg.in/mgo.v2"
)

func MongoConnect() *mgo.Session {
	var sessionDb, err = mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err.Error())
	}

	return sessionDb
}
