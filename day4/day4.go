package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/rosswf/aoc2022/utils"
)

//go:embed input.txt
var input string

func main() {
	pairings := parse(strings.TrimSuffix(input, "\n"))
	p1 := part1(pairings)
	p2 := part2(pairings)
	fmt.Printf("Day 4-Part 1: %d\n", p1)
	fmt.Printf("Day 4-Part 2: %d\n", p2)
}

func part1(p *[][]int) int {
	contained := 0
	for _, v := range *p {
		if (v[0] <= v[2] && v[3] <= v[1]) || (v[2] <= v[0] && v[1] <= v[3]) {
			contained++
		}
	}
	return contained
}

func part2(p *[][]int) int {
	overlap := 0
	for _, v := range *p {
		if v[0] <= v[3] && v[1] >= v[2] {
			overlap++
		}
	}
	return overlap
}

func parse(s string) *[][]int {
	ss := strings.Split(s, "\n")
	pairings := make([][]int, len(ss))
	for i, v := range strings.Split(s, "\n") {
		p := strings.ReplaceAll(v, ",", "-")
		pairings[i] = utils.Map(strings.Split(p, "-"), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
	}
	return &pairings
}
