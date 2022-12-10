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
	c := parse(ti)
	got := part1(c)
	want := 13140

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
