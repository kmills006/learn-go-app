package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

// League represents a list of players
type League []Player

// Find will return an individual player by their name
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

// NewLeague returns the decoded league from the JSON data store
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
