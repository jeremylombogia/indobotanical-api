package config

import (
	"gopkg.in/mgo.v2"
)

func MongoConnect() (*mgo.Session, error) {
	var session, err = mgo.Dial("mongodb://heroku_thr3t37f:d9qqnoegat1lam2dpkop0sekjm@ds235431.mlab.com:35431/heroku_thr3t37f")
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	return session, nil
}
