package poker

import "testing"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

// GetPlayerScore stubs PlayerStore GetPlayerScore
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

// RecordWin stubs PlayerStore RecordWin
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// GetLeague stubs PlayerStore GetLeague
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// AssertPlayerWin validates that the expected winner is returned
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %s, want %s", store.winCalls[0], winner)
	}
}
