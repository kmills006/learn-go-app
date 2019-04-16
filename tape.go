package main

import "os"

// Separate out the concerns of the kind of data we write, from the writing
// Encapsulate out "when we write we go from the beginning functionality"
type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)

	return t.file.Write(p)
}
