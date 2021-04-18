package main

import (
	"log"
	"net/http"
	"strings"
)

type InMemoryPlayerStore struct{
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore{
	return &InMemoryPlayerStore{map[string]int{}}
}


func  (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	name = strings.ToLower(name)
	if name == "pepper" {
		return 20
	}

	if name == "floyd" {
		return 10
	}

	return 0
}

func  (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
