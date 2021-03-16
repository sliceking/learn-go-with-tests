package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountdownOperationsSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("sleep before we print", func(t *testing.T) {
		spySleeper := &CountdownOperationsSpy{}
		Countdown(spySleeper, spySleeper)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
		}
	})

	t.Run("has correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &CountdownOperationsSpy{}

		Countdown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if len(spy.Calls) != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", len(spy.Calls))
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 *time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
