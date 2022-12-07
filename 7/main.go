package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const AvailableSpace = 70000000

type Node struct {
	id       string
	size     int
	children []*Node
	parent   *Node
}

func ReadLines() []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func CreateDirectoryHierarchy(lines []string) Node {
	root := Node{
		id:       "/",
		size:     0,
		children: make([]*Node, 0),
		parent:   nil,
	}

	curr := &root
	parent := &root.parent
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					curr = *parent
					parent = &curr.parent
				} else if tokens[2] != "/" {
					nextParent := curr
					nextCurr := Node{
						id:       tokens[2],
						size:     0,
						children: make([]*Node, 0),
						parent:   nextParent,
					}
					curr.children = append(curr.children, &nextCurr)
					curr = &nextCurr
					parent = &nextParent
				}

			}
		} else if tokens[0] != "dir" {
			bytes, _ := strconv.Atoi(tokens[0])
			curr.size += bytes

			parent := curr.parent
			for parent != nil {
				parent.size += bytes
				parent = parent.parent
			}
		}
	}

	return root
}

func FindLargeDirectories(root Node, limit int) []Node {
	dirs := make([]Node, 0)

	if root.size <= limit {
		dirs = append(dirs, root)
	}

	for _, child := range root.children {
		dirs = append(dirs, FindLargeDirectories(*child, limit)...)
	}

	return dirs
}

func SumDirSizes(dirs []Node) (size int) {
	for _, dir := range dirs {
		size += dir.size
	}
	return size
}

func FindDirToDelete(dirs []Node, unused int, need int) (Node, bool) {
	smallest := Node{
		size: math.MaxInt,
	}

	if unused >= need {
		return smallest, false
	}

	for _, dir := range dirs {
		if unused+dir.size >= need {
			if dir.size < smallest.size {
				smallest = dir
			}
		}
	}
	return smallest, false
}

func main() {
	lines := ReadLines()
	root := CreateDirectoryHierarchy(lines)

	// Part 1
	dirs := FindLargeDirectories(root, 100000)
	fmt.Println(SumDirSizes(dirs))

	// Part 2
	dirs = FindLargeDirectories(root, math.MaxInt)
	unusedSpace := AvailableSpace - root.size
	smallest, _ := FindDirToDelete(dirs, unusedSpace, 30000000)
	fmt.Println(smallest.size)
}
