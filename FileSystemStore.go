package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// FileSystemPlayerStore stores players in the data store
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore handles the initialization of the File System data store
func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	database.Seek(0, 0)

	// file.Stat returns stats on our file. This lets us check the size of the file,
	// if it's empty we Write an empty JSON array and Seek back to the start
	// ready for the rest of the code.
	info, err := database.Stat()

	if err != nil {
		return nil, fmt.Errorf("problem getting file info from file %s, %v", database.Name(), err)
	}

	if info.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}

	league, err := NewLeague(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{database}),
		league:   league,
	}, nil
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

	f.database.Encode(f.league)
}
