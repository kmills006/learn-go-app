package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemPlayerStore stores players in the data store
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// initializePlayerDbFile initializes the store DB file
func initializePlayerDbFile(file *os.File) error {
	file.Seek(0, 0)

	// file.Stat returns stats on our file. This lets us check the size of the file,
	// if it's empty we Write an empty JSON array and Seek back to the start
	// ready for the rest of the code.
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

// NewFileSystemPlayerStore handles the initialization of the File System data store
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDbFile(file)

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

// GetLeague returns the players in the league from the data store
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

// GetPlayerStore returns an individual players score
func (f *FileSystemPlayerStore) GetPlayerStore(name string) int {
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
