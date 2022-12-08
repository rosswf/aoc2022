package main

import (
	_ "embed"
	"reflect"
	"strings"
	"testing"
)

//go:embed test-input.txt
var ti string

func TestParse(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	got := parse(ti)
	want := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	if !reflect.DeepEqual(*got.trees, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	trees := parse(ti)
	got := part1(trees)
	want := 21

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	ti := strings.TrimSuffix(ti, "\n")
	trees := parse(ti)
	got := part2(trees)
	want := 8

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
