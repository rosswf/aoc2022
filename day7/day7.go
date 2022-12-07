package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var s string

type Node struct {
	name     string
	size     int // 0 if dir
	parent   *Node
	children map[string]*Node
}

func init() {
	s = strings.TrimSuffix(s, "\n")
}

func main() {
	root := FS(s)
	calcDirSizes(root)

	// Part 1
	total := new(int)
	limit := 100000
	sumChildren(root, limit, total)
	fmt.Printf("Day 7-Part 1: %d\n", *total)

	// Part 2
	requiredSpace := root.size - 40000000
	dirSizes := []int{}
	dirsOverRequired(root, &dirSizes, requiredSpace)
	sort.Ints(dirSizes)
	fmt.Printf("Day 7-Part 2: %d\n", dirSizes[0])
}

func dirsOverRequired(n *Node, dirSizes *[]int, required int) {
	for _, child := range n.children {
		if child.children != nil {
			if child.size >= required {
				*dirSizes = append(*dirSizes, child.size)
			}
		}
		dirsOverRequired(child, dirSizes, required)
	}
}

func calcDirSizes(n *Node) {
	if n.children != nil {
		for _, child := range n.children {
			calcDirSizes(child)
			n.size += child.size
		}
	}
}

func sumChildren(n *Node, limit int, acc *int) {
	for _, child := range n.children {
		if child.children != nil {
			if child.size <= limit {
				*acc += child.size
			}
		}
		sumChildren(child, limit, acc)
	}
}

func FS(s string) *Node {
	root := &Node{
		size:     0,
		parent:   nil,
		children: map[string]*Node{},
	}

	commands := strings.Split(s, "\n")

	current_node := root

	for _, command := range commands[1:] {
		if command == "$ ls" {
			continue
		}
		if strings.HasPrefix(command, "$ cd ") {
			dir := strings.TrimPrefix(command, "$ cd ")
			if dir == ".." {
				current_node = current_node.parent
			} else {
				current_node = current_node.children[dir]
			}
		} else {
			file_info := strings.Split(command, " ")
			size := file_info[0]
			name := file_info[1]
			if size == "dir" {
				current_node.children[name] = &Node{
					name:     name,
					size:     0,
					parent:   current_node,
					children: map[string]*Node{},
				}
			} else {
				size, _ := strconv.Atoi(size)
				current_node.children[name] = &Node{
					name:     name,
					size:     size,
					parent:   current_node,
					children: nil,
				}
			}
		}
	}
	return root
}
