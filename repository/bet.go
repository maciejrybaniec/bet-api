package repository

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Bet - structure for bet model
type Bet struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	EventID string        `json:"eventId"`
}

// Bets - collection of bet models
type Bets []Bet

// Save - saves model in database
func (b Bet) Save() error {
	session, connection := MongoConnection()
	defer session.Close()

	collection := connection.C("bets")
	err := collection.Insert(b)
	if err != nil {
		log.Fatal(err)
		return errors.New("Cannot save bet model")
	}

	return nil
}
