package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// CLI : TODO
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// NewCLI is a constructor that will automatically wrap the Scanner
func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		alerter:     alerter,
		in:          bufio.NewScanner(in),
		playerStore: store,
	}
}

// PlayPoker starts a new game of poker from the CLI
func (c *CLI) PlayPoker() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}

	c.alerter.ScheduleAlertAt(5*time.Second, 100)

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
