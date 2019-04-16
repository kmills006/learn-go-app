package main

import (
	"io/ioutil"
	"testing"
)

func TestTape(t *testing.T) {
	file, cleanDatabase := createTempFile(t, "12345")

	defer cleanDatabase()

	tape := &tape{file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)

	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
