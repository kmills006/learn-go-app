package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// StubPlayerStore implements PlayerStore for testing purposes
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

// type PlayerStore interface {
// }

// GetPlayerScore returns a score from Scores
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin will record a win to WinCalls
func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// GetLeague returns League
func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

// GetLeagueFromResponse parses the response from GET: /league
func GetLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return
}

// NewGetScoreRequest creates a new GET request to /players/{Name}
func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

// NewGetLeagueRequest creates a new GET request to /players/{Name}
func NewGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)

	return req
}

// NewPostWinRequest creates a new POST request to /players/{Name}
func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return req
}

// AssertContentType checks the content type of the response
func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}

}

// AssertLeague checks that the expected league matches the returned league
func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// AssertStatus checks that the server responded with the expected HTTP status code
func AssertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// AssertResponseBody checks that the response body matches the returned response body
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %s, want %s", got, want)
	}
}

// AssertPlayerWin checks that the expected winner is returned
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %s, want %s", store.WinCalls[0], winner)
	}
}

// AssertNoError checks if an error was returned
func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

// AssertScoreEquals compares the expected score vs the actual score
func AssertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
