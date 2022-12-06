package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var s string

var Empty struct{}

func main() {
	p1 := findMarker(s, 4)
	p2 := findMarker(s, 14)

	fmt.Printf("Day 6-Part 1: %d\n", p1)
	fmt.Printf("Day 6-Part 2: %d\n", p2)

}

func findMarker(s string, n int) int {
	for i := 0; i < len(s)-n; i++ {
		marker := make(map[byte]struct{})
		for j := 0; j < n; j++ {
			marker[s[i+j]] = Empty
		}
		if len(marker) == n {
			return i + n
		}
	}
	return -1
}
