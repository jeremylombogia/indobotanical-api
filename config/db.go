package config

import (
	"os"

	"gopkg.in/mgo.v2"
)

func MongoConnect() (*mgo.Session, error) {
	var session, err = mgo.Dial(os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	return session, nil
}
