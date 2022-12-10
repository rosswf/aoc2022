package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var s string

func main() {
	s := strings.TrimSuffix(s, "\n")
	commands := parse(s)
	p1 := part1(commands)

	fmt.Printf("Day 10-Part 1: %d\n", p1)
}

func part1(commands *[]int) int {
	strengths := 0
	strength := 1
	cycle := 0
	for _, c := range *commands {
		cycle++
		if (cycle-20)%40 == 0 {
			strengths += cycle * strength
		}
		if c != 0 {
			cycle++
			if (cycle-20)%40 == 0 {
				strengths += cycle * strength
			}
			strength += c
		}
	}
	return strengths
}

func parse(s string) *[]int {
	commands := []int{}
	for _, c := range strings.Split(s, "\n") {
		if strings.HasPrefix(c, "addx") {
			i, _ := strconv.Atoi(strings.SplitAfter(c, " ")[1])
			commands = append(commands, i)

		} else {
			commands = append(commands, 0)
		}
	}
	return &commands
}
