package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var s string

type forest struct {
	trees *[][]int
	rows  int
	cols  int
}

func main() {
	s := strings.TrimSuffix(s, "\n")
	f := parse(s)
	p1 := part1(f)
	p2 := part2(f)

	fmt.Printf("Day 8-Part 1: %d\n", p1)
	fmt.Printf("Day 8-Part 2: %d\n", p2)
}

func part2(forest *forest) int {
	highest := 0
	for y := 1; y < forest.rows-1; y++ {
		for x := 1; x < forest.cols-1; x++ {
			north, south, east, west := checkTree(x, y, (*forest.trees)[y][x], forest)
			if north != y {
				north++
			}
			if south != forest.cols-1-y {
				south++
			}
			if west != x {
				west++
			}
			if east != forest.rows-1-x {
				east++
			}

			scenic := north * south * east * west

			if scenic > highest {
				highest = scenic
			}
		}
	}
	return highest
}

func part1(forest *forest) int {
	visible := forest.cols*2 + forest.rows*2 - 4
	for y := 1; y < forest.rows-1; y++ {
		for x := 1; x < forest.cols-1; x++ {
			north, south, east, west := checkTree(x, y, (*forest.trees)[y][x], forest)
			if west == x || forest.rows-1-east == x || north == y || forest.cols-1-south == y {
				visible++
			}
		}
	}
	return visible
}

func checkTree(x, y, currentTree int, forest *forest) (int, int, int, int) {
	west := 0
	for i := x - 1; i >= 0; i-- {
		if currentTree > (*forest.trees)[y][i] {
			west++
			continue
		}
		break
	}

	east := 0
	for i := x + 1; i < forest.rows; i++ {
		if currentTree > (*forest.trees)[y][i] {
			east++
			continue
		}
		break
	}

	north := 0
	for i := y - 1; i >= 0; i-- {
		if currentTree > (*forest.trees)[i][x] {
			north++
			continue
		}
		break
	}

	south := 0
	for i := y + 1; i < forest.cols; i++ {
		if currentTree > (*forest.trees)[i][x] {
			south++
			continue
		}
		break
	}
	return north, south, east, west
}

func parse(s string) *forest {
	rows := strings.Split(s, "\n")
	cols := len(rows[0])
	trees := make([][]int, cols)

	for x, row := range rows {
		trees[x] = make([]int, len(row))
		for y, tree := range row {
			tree, _ := strconv.Atoi(string(tree))
			trees[x][y] = tree
		}
	}
	return &forest{
		trees: &trees,
		rows:  len(rows),
		cols:  cols,
	}
}
