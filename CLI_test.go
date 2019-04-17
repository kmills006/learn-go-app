package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/learn-go-app"
)

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		var dummySpyAlert = &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlert)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		var dummySpyAlert = &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlert)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []struct {
			expectedScheduleTime time.Duration
			expectedAmount       int
		}{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			var description = fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime)

			t.Run(description, func(t *testing.T) {
				if len(blindAlerter.alerts) < 1 {
					t.Fatalf("alert %d was not scheduled for %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]
				amountGot := alert.amount

				if amountGot != c.expectedAmount {
					t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
				}

				scheduledTimeGot := alert.at

				if scheduledTimeGot != c.expectedScheduleTime {
					t.Errorf("got scheduled time of %v, want %v", scheduledTimeGot, c.expectedScheduleTime)
				}
			})
		}
	})
}
