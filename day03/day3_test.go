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
	rucksacks := strings.Split(ti, "\n")
	got := part1(rucksacks)
	want := 157

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	ti = strings.TrimSuffix(ti, "\n")
	rucksacks := strings.Split(ti, "\n")
	got := part2(rucksacks)
	want := 70

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
