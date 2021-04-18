package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) ClearWinRecord() {
	s.winCalls = []string{}
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"pepper": 20,
			"floyd": 10,
		},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		store.ClearWinRecord()
		request := newPostWinRequest("pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins on POST", func(t *testing.T) {
		store.ClearWinRecord()

		player := "pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
