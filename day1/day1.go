package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	data := parseInput(input)
	sort.Sort(sort.Reverse(sort.IntSlice(data)))

	p1 := part1(data)
	fmt.Printf("Day 1-Part 1: %d\n", p1)

	p2 := part2(data)
	fmt.Printf("Day 1-Part 2: %d\n", p2)
}

func part1(i []int) int {
	return i[0]
}

func part2(i []int) int {
	return sum(i[:3])
}

func parseInput(s string) []int {
	var (
		totals []int
		total  int
	)

	split := strings.Split(s, "\n")
	for _, v := range split {
		if v == "" {
			totals = append(totals, total)
			total = 0
			continue
		}
		v, _ := strconv.Atoi(v)
		total += v
	}
	totals = append(totals, total)

	return totals
}

func sum(ints []int) (sum int) {
	for _, v := range ints {
		sum += v
	}
	return
}
