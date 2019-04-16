package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 23}]`)

		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 23},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 23}]`)

		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 23

		assertScoreEquals(t, got, want)
	})
}

// createTempFile creates a temporary file with some data inside of it
func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	// TempFile creates a temporary file
	// "db" is a prefix to put on the random file created
	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()

		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

// assertScoreEquals compares the expected score vs the actual score
func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
