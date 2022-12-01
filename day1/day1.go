package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var s string

func main() {
	d := parseInput(s)

	p1 := part1(d)
	fmt.Printf("Day 1-Part 1: %d\n", p1)

	p2 := part2(d)
	fmt.Printf("Day 1-Part 2: %d\n", p2)
}

func part1(i []int) int {
	return i[0]
}

func part2(i []int) int {
	return sum(i[:3])
}

func parseInput(s string) []int {
	var totals []int
	var t int

	split := strings.Split(s, "\n")
	for _, v := range split {
		if v == "" {
			totals = append(totals, t)
			t = 0
			continue
		}
		cal, _ := strconv.Atoi(v)
		t += cal
	}
	totals = append(totals, t)

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	return totals
}

func sum(ints []int) (sum int) {
	for _, v := range ints {
		sum += v
	}
	return
}
