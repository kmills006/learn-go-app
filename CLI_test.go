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

	if len(playerScore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := playerScore.winCalls[0]
	want := "Chris"

	if got != want {
		t.Errorf("didn't record correct winner, got %s, want %s", got, want)
	}
}
