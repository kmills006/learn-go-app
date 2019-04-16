package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerScore := &StubPlayerScore{}

		cli := &CLI{playerScore, in}

		cli.PlayPoker()

		assertPlayerWin(t, playerScore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerScore := &StubPlayerScore{}

		cli := &CLI{playerScore, in}

		cli.PlayPoker()

		assertPlayerWin(t, playerScore, "Cleo")
	})
}
