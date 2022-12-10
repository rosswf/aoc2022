package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input string

	shapes = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	shapePointValues = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	winValues = map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	loseValues = map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}
)

type round struct {
	opponent string
	player   string
}

func init() {
	input = strings.TrimSuffix(input, "\n")
}

func main() {
	rounds := parseInput(input)
	p1 := part1(rounds)
	p2 := part2(rounds)
	fmt.Printf("Day 2-Part 1: %d\n", p1)
	fmt.Printf("Day 2-Part 2: %d\n", p2)
}

func part1(r *[]round) int {
	shapePoints := calculateShapePoints(r, convertPlayerMove)
	roundPoints := calculateRoundPoints(r)
	return shapePoints + roundPoints
}

func part2(r *[]round) int {
	shapePoints := calculateShapePoints(r, calculatePlayerMove)
	roundPoints := calculateRoundPointsP2(r)
	return shapePoints + roundPoints
}

func parseInput(s string) *[]round {
	rounds := []round{}
	for _, v := range strings.Split(s, "\n") {
		r := strings.Split(v, " ")
		rounds = append(rounds, round{
			opponent: r[0],
			player:   r[1],
		})
	}
	return &rounds
}

func calculateShapePoints(r *[]round, f func(r round) string) int {
	p := 0
	for _, round := range *r {
		m := f(round)
		p += shapePointValues[m]
	}
	return p
}

func calculateRoundPoints(r *[]round) int {
	var p int
	for _, round := range *r {
		switch shapes[round.player] {
		case winValues[round.opponent]:
			p += 6
		case round.opponent:
			p += 3
		}
	}
	return p
}

func calculateRoundPointsP2(r *[]round) int {
	var p int
	for _, round := range *r {
		switch round.player {
		case "Y":
			p += 3
		case "Z":
			p += 6
		}
	}
	return p
}

func convertPlayerMove(r round) string {
	return shapes[r.player]
}

func calculatePlayerMove(r round) string {
	switch r.player {
	// Lose
	case "X":
		return loseValues[r.opponent]
	// Draw
	case "Y":
		return r.opponent
	// Win
	default:
		return winValues[r.opponent]
	}
}
