package main

import (
	tdd_kata "github.com/aamir-form3/tdd-kata"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &tdd_kata.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5050", server))
}
