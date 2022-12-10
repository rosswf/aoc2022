package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test-input.txt
var ti string

func TestPart1(t *testing.T) {
	ti = strings.TrimSuffix(ti, "\n")
	parsed := parse(ti)
	got := part1(parsed)
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {

	ti = strings.TrimSuffix(ti, "\n")
	parsed := parse(ti)
	got := part2(parsed)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
