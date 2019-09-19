package config

import (
	"gopkg.in/mgo.v2"
)

func MongoConnect() (*mgo.Session, error) {
	var session, err = mgo.Dial("mongodb://heroku_ghdkzt3f:61o017lubhkdmddl2qc2jkiko1@ds137090.mlab.com:37090/heroku_ghdkzt3f")
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	return session, nil
}
