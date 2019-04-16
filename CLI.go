package poker

import "io"

// CLI : TODO
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

// PlayPoker starts a new game of poker from the CLI
func (c *CLI) PlayPoker() {
	c.playerStore.RecordWin("Chris")
}
