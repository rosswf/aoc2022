package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test-input.txt
var ti string

func TestPart1(t *testing.T) {
	split := strings.Split(strings.TrimSuffix(ti, "\n"), "\n\n")
	st := split[0]
	m := split[1]
	stacks := *parseStacks(st)
	moves := *parseMoves(m)

	got := part1(stacks, moves)
	want := "CMZ"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	split := strings.Split(strings.TrimSuffix(ti, "\n"), "\n\n")
	st := split[0]
	m := split[1]
	stacks := *parseStacks(st)
	moves := *parseMoves(m)

	got := part2(stacks, moves)
	want := "MCD"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
