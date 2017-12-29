package main

import (
	"encoding/json"
	"log"
	"net/http"
	"server/repository"
)

const port string = ":8000"

func placeBetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	if r.Body == nil {
		http.Error(w, "Request body is required", 400)
		return
	}

	var bet repository.Bet
	if err := json.NewDecoder(r.Body).Decode(&bet); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := bet.Save(); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
}

func betsListHandler(w http.ResponseWriter, r *http.Request) {
	session, connection := repository.MongoConnection()
	defer session.Close()

	collection := connection.C("bets")

	var bets repository.Bets
	err := collection.Find(nil).All(&bets)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bets)
}

func main() {
	log.Println("Server is running on port", port)

	http.HandleFunc("/bets", betsListHandler)
	http.HandleFunc("/bet", placeBetHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
