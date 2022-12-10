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
	fmt.Println("Day 10-Part 2:")
	part2(commands)
}

func part2(commands *[]int) {
	X := 1
	cycle := 0
	row := ""
	for _, c := range *commands {
		row = draw(cycle, X, row)
		cycle++
		if c != 0 {
			row = draw(cycle, X, row)
			cycle++
			X += c
		}
	}
}

func draw(cycle, X int, row string) string {
	switch cycle % 40 {
	case X - 1, X, X + 1:
		row += "#"
	default:
		row += "."
	}
	if len(row)%40 == 0 {
		fmt.Println(row)
		row = ""
	}
	return row
}

func part1(commands *[]int) int {
	total := 0
	X := 1
	cycle := 0
	for _, c := range *commands {
		cycle++
		if (cycle-20)%40 == 0 {
			total += cycle * X
		}
		if c != 0 {
			cycle++
			if (cycle-20)%40 == 0 {
				total += cycle * X
			}
			X += c
		}
	}
	return total
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
