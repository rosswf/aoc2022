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
	rounds := parseInput(ti)
	got := part1(rounds)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	ti = strings.TrimSuffix(ti, "\n")
	rounds := parseInput(ti)
	got := part2(rounds)
	want := 12

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
