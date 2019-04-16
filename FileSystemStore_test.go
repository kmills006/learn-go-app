package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 23}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 23},
		}

		assertLeague(t, got, want)
	})
}
