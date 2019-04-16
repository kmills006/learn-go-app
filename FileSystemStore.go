package main

import (
	"io"
)

// FileSystemPlayerStore stores players in the data store
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague returns the players in the league from the data store
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)

	league, _ := NewLeague(f.database)

	return league
}

// GetPlayerScore returns an individual players score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins

			break
		}
	}

	return wins
}
