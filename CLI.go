package poker

// CLI : TODO
type CLI struct {
	playerScore PlayerScore
}

// PlayPoker starts a new game of poker from the CLI
func (c *CLI) PlayPoker() {
	c.playerScore.RecordWin("Cleo")
}
