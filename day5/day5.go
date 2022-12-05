package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rosswf/aoc2022/utils"
)

type Stack []string

func (st *Stack) Push(s string) {
	*st = append(*st, s)
}

func (st *Stack) Pop() string {
	i := len(*st) - 1
	v := (*st)[i]
	*st = (*st)[:i]
	return v
}

type move struct {
	qty  int
	from int
	to   int
}

//go:embed input.txt
var s string

func main() {
	split := strings.Split(strings.TrimSuffix(s, "\n"), "\n\n")
	st := split[0]
	m := split[1]

	moves := *parseMoves(m)

	stacks := *parseStacks(st)
	p1 := part1(stacks, moves)

	stacks = *parseStacks(st)
	p2 := part2(stacks, moves)

	fmt.Printf("Day 5-Part 1: %s\n", p1)
	fmt.Printf("Day 5-Part 2: %s\n", p2)
}

func part1(s []Stack, moves []move) string {
	for _, m := range moves {
		for i := 0; i < m.qty; i++ {
			v := s[m.from-1].Pop()
			s[m.to-1].Push(v)
		}
	}

	var sb strings.Builder

	for _, stack := range s {
		fmt.Fprintf(&sb, "%s", stack.Pop())
	}

	return sb.String()
}

func part2(s []Stack, moves []move) string {
	for _, m := range moves {
		picked := make(Stack, m.qty)
		for i := m.qty - 1; i >= 0; i-- {
			picked[i] = s[m.from-1].Pop()
		}
		s[m.to-1] = append(s[m.to-1], picked...)
	}

	var sb strings.Builder

	for _, stack := range s {
		fmt.Fprintf(&sb, "%s", stack.Pop())
	}

	return sb.String()
}

func parseStacks(s string) *[]Stack {
	tallest := strings.Count(s, "\n")
	qtyStacks, _ := strconv.Atoi(string(s[len(s)-2]))

	stacks := make([]Stack, qtyStacks)

	for i, stack := range strings.Split(s, "\n") {
		stack = "  " + stack
		for j, c := range stack {
			if c == '[' {
				stackNo := (j+2)/4 - 1
				crate := stack[j+1]

				if stacks[stackNo] == nil {
					stacks[stackNo] = make(Stack, tallest-i)
				}

				stacks[stackNo][tallest-i-1] = string(crate)
			}
		}
	}
	return &stacks
}

func parseMoves(s string) *[]move {
	moves := []move{}

	r, _ := regexp.Compile(`\d+`)

	for _, v := range strings.Split(s, "\n") {
		matches := utils.Map(r.FindAllString(v, -1), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		moves = append(moves,
			move{
				qty:  matches[0],
				from: matches[1],
				to:   matches[2],
			})
	}

	return &moves
}
