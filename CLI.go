package poker

import "io"

// CLI : TODO
type CLI struct {
	playerScore PlayerScore
	in          io.Reader
}

// PlayPoker starts a new game of poker from the CLI
func (c *CLI) PlayPoker() {
	c.playerScore.RecordWin("Chris")
}
