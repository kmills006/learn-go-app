package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore : TODO
type InMemoryPlayerStore struct{}

// GetPlayerScore : TODO
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
