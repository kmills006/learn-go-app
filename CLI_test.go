package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	playerScore := &StubPlayerScore{}

	cli := &CLI{playerScore, in}

	cli.PlayPoker()

	assertPlayerWin(t, playerScore, "Chris")
}
