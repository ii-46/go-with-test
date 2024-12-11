package main

import (
	"bytes"
	"reflect"
	"testing"
)

const sleep = "sleep"
const write = "write"

type SpyCountDownOperation struct {
	Calls []string
}

func (s *SpyCountDownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpySleeper struct {
	Call int
}

func (s *SpySleeper) Sleep() {
	s.Call++
}

func TestCount(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
	if spySleeper.Call != 4 {
		t.Errorf("want 4 got %d", spySleeper.Call)
	}
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperation{}
		Countdown(spySleepPrinter, spySleepPrinter)
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
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter)
		}
	})
}
