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

type StubPlayerScore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerScore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerScore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerScore) GetLeague() League {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerScore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerScore{
		map[string]int{},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %s, want %s", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerScore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 10},
			{"Tiesta", 15},
		}

		store := StubPlayerScore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertContentType(t, response, "application/json")
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func newGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)

	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}

}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// assertStatus : Validate that the server responded with the expected HTTP status code
func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %s, want %s", got, want)
	}
}

func assertPlayerWin(t *testing.T, store *StubPlayerScore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], winner)
	}
}
