package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database := strings.NewReader(`[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 23}]`)

	t.Run("/league from a reader", func(t *testing.T) {
		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 23},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 23

		assertScoreEquals(t, got, want)
	})
}

// assertScoreEquals compares the expected score vs the actual score
func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
