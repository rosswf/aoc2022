package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test-input.txt
var ti string

//go:embed test-input2.txt
var ti2 string

func TestPart1(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	got := part1(ti)
	want := 13

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestPart2(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	got := part2(ti)
	want := 1

	ti2 := strings.TrimSuffix(ti2, "\n")
	got2 := part2(ti2)
	want2 := 36

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	if got != want {
		t.Errorf("got %d want %d", got2, want2)
	}
}
