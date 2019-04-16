package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore stores players in the data store
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

// NewFileSystemPlayerStore handles the initialization of the File System data store
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)

	league, _ := NewLeague(database)

	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}

// GetLeague returns the players in the league from the data store
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

// GetPlayerScore returns an individual players score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increments the players score by 1
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0, 0)

	json.NewEncoder(f.database).Encode(f.league)
}
