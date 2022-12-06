package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test-input.txt
var ti string

func TestPart1(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	got := findMarker(ti, 4)

	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	got := findMarker(ti, 14)

	want := 19

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
