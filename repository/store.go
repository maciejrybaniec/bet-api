package repository

import (
	mgo "gopkg.in/mgo.v2"
)

// MongoConnection - establish connection with database
func MongoConnection() (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	connection := session.DB("bets-app")

	return session, connection
}
