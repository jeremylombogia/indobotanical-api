package config

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func MongoConnect() *mgo.Session {
	host := fmt.Sprintf("%s", os.Getenv("DB_HOST"))
	info := &mgo.DialInfo{
		Addrs:    []string{host},
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	var sessionDb, err = mgo.DialWithInfo(info)
	if err != nil {
		panic(err.Error())
	}

	return sessionDb
}
