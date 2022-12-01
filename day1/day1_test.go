package main

import (
	_ "embed"
	"testing"
)

//go:embed test-input.txt
var ts string

func TestPart1(t *testing.T) {
	i := parseInput(ts)
	got := part1(i)
	want := 24000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	i := parseInput(ts)
	got := part2(i)
	want := 45000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
