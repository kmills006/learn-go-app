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
func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)

	league, _ := NewLeague(f.database)

	return league
}

// GetPlayerScore returns an individual players score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increments the players score by 1
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)

	json.NewEncoder(f.database).Encode(league)
}
