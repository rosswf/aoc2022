package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Stack []string

//go:embed input.txt
var s string

func main() {
	split := strings.Split(strings.TrimSuffix(s, "\n"), "\n\n")
	st := split[0]
	moves := split[1]
	stacks := *parseStacks(st)

	fmt.Printf("%v\n", stacks)
	fmt.Println(moves)
}

func parseStacks(s string) *[]Stack {

	tallest := strings.Count(s, "\n")

	numberStacks, _ := strconv.Atoi(string(s[len(s)-2]))

	stacks := make([]Stack, numberStacks)

	for i, x := range strings.Split(s, "\n") {
		x = "  " + x
		for j, y := range x {
			if y == '[' {
				stackNo := (j+2)/4 - 1
				crate := x[j+1]

				if stacks[stackNo] == nil {
					stacks[stackNo] = make(Stack, tallest-i)
				}

				stacks[stackNo][tallest-i-1] = string(crate)
			}
		}
	}
	return &stacks
}

//
//
//
//
//
//
//
//
//
