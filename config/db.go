package config

import (
	"gopkg.in/mgo.v2"
)

func MongoConnect() (*mgo.Session, error) {
	var session, err = mgo.Dial("mongodb+srv://jeremy:jeremi11@indobotanical-t4rkd.mongodb.net/test?retryWrites=true&w=majority")
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	return session, nil
}
