package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/rosswf/aoc2022/utils"
)

//go:embed input.txt
var s string

type command struct {
	direction string
	distance  int
}

var Empty struct{}

type point struct {
	x       int
	y       int
	visited map[[2]int]struct{}
}

func init() {
	s = strings.TrimSuffix(s, "\n")
}

func main() {
	commands := parse(s)
	p1 := part1(commands)
	p2 := part2(commands)

	fmt.Printf("Day 9-Part 1: %d\n", p1)
	fmt.Printf("Day 9-Part 2: %d\n", p2)
}

func part1(commands *[]command) int {
	h := &point{
		x: 0,
		y: 0,
	}

	t0 := &point{
		x: 0,
		y: 0,
		visited: map[[2]int]struct{}{
			{0, 0}: Empty,
		},
	}

	tails := []*point{t0}

	moveHead(commands, h, tails)

	return len(tails[0].visited)
}

func part2(commands *[]command) int {
	h := &point{
		x: 0,
		y: 0,
	}

	tails := []*point{}

	for i := 0; i < 9; i++ {
		tails = append(tails,
			&point{
				x: 0,
				y: 0,
				visited: map[[2]int]struct{}{
					{0, 0}: Empty,
				},
			})
	}

	moveHead(commands, h, tails)

	return len(tails[8].visited)
}

func moveHead(commands *[]command, head *point, tails []*point) {
	for _, c := range *commands {
		for i := 0; i < c.distance; i++ {
			h := head
			switch c.direction {
			case "U":
				h.y++
			case "D":
				h.y--
			case "L":
				h.x--
			case "R":
				h.x++
			}
			for _, t := range tails {
				if utils.Abs(h.x-t.x) > 1 || utils.Abs(h.y-t.y) > 1 {
					moveTail(t, h, c.direction)
				}
				h = t
			}
		}
	}
}

func moveTail(t, h *point, d string) {
	// diagonal
	if t.x != h.x && t.y != h.y {
		switch h.y - t.y {
		case 2, 1:
			t.y++
		case -2, -1:
			t.y--
		}
		switch h.x - t.x {
		case -2, -1:
			t.x--
		case 2, 1:
			t.x++
		}
	} else {
		switch h.y - t.y {
		case 2:
			t.y++
		case -2:
			t.y--
		}
		switch h.x - t.x {
		case -2:
			t.x--
		case 2:
			t.x++
		}
	}
	t.visited[[2]int{t.x, t.y}] = Empty
}

func parse(s string) *[]command {
	commands := []command{}
	for _, i := range strings.Split(s, "\n") {
		c := strings.Split(i, " ")
		d, _ := strconv.Atoi(c[1])
		commands = append(commands,
			command{
				direction: c[0],
				distance:  d,
			})
	}
	return &commands
}
