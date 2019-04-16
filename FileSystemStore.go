package main

import (
	"io"
)

// FileSystemPlayerStore stores players in the data store
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague returns the players in the league from the data store
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)

	league, _ := NewLeague(f.database)

	return league
}
