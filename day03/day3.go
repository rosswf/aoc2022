package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimSuffix(input, "\n")
}

func main() {
	elves := strings.Split(input, "\n")
	p1 := part1(elves)
	p2 := part2(elves)
	fmt.Printf("Day 3-Part 1: %d\n", p1)
	fmt.Printf("Day 3-Part 2: %d\n", p2)
}

func part1(elves []string) (p int) {
	for _, e := range elves {
		firstRucksack := e[:len(e)/2]
		secondRucksack := e[len(e)/2:]

		for _, s := range firstRucksack {
			if strings.Contains(secondRucksack, string(s)) {
				p += priority(byte(s))
				break
			}
		}
	}
	return
}

func part2(elves []string) (p int) {
	for i := 0; i < len(elves); i += 3 {
		for _, s := range elves[i] {
			if strings.Contains(elves[i+1], string(s)) && strings.Contains(elves[i+2], string(s)) {
				p += priority(byte(s))
				break
			}
		}
	}
	return
}

func priority(c byte) int {
	if c > 90 {
		return int(c) - 96
	}
	return int(c) - 64 + 26
}
