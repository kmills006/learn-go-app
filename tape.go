package main

import "io"

// Separate out the concerns of the kind of data we write, from the writing
// Encapsulate out "when we write we go from the beginning functionality"
type tape struct {
	file io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Seek(0, 0)

	return t.file.Write(p)
}
