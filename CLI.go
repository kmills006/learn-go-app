package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI : TODO
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

// NewCLI is a constructor that will automatically wrap the Scanner
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

// PlayPoker starts a new game of poker from the CLI
func (c *CLI) PlayPoker() {
	userInput := c.readLine()

	c.playerStore.RecordWin(extractWinner(userInput))
}

// extractWinner
func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()

	return c.in.Text()
}
