package main

import (
	_ "embed"
	"sort"
	"testing"
)

//go:embed test-input.txt
var ti string

func TestPart1(t *testing.T) {
	d := parseInput(ti)
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	got := part1(d)
	want := 24000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	d := parseInput(ti)
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	got := part2(d)
	want := 45000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
