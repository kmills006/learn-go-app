package main

// NewInMemoryPlayerStore initializes an empty player store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore collects data about players in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetLeague returns the plays stored in memory
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}

// RecordWin will record a players win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore retrieves scores for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
