package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines() []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func AreAllDifferentChars(s string) bool {
	tracker := make(map[rune]bool)
	for _, c := range s {
		_, exists := tracker[c]
		if exists {
			return false
		}
		tracker[c] = true
	}
	return true
}

func FindStartOfMarker(signal string, n int) int {
	idx := -1

	for i := 0; i < len(signal)-n; i++ {
		if AreAllDifferentChars(signal[i : i+n]) {
			return i + n
		}
	}

	return idx
}

func main() {
	lines := ReadLines()
	signal := lines[0]

	// Part 1
	fmt.Println(FindStartOfMarker(signal, 4))

	// Part 2
	fmt.Println(FindStartOfMarker(signal, 14))
}
