package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI : TODO
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

// PlayPoker starts a new game of poker from the CLI
func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()

	c.playerStore.RecordWin(extractWinner(reader.Text()))
}

// extractWinner
func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
