package main

import (
	"encoding/json"
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

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)

	json.NewEncoder(f.database).Encode(league)
}
