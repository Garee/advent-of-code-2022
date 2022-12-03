package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func FindSharedItem(first string, second string) (rune, bool) {
	for _, c := range first {
		if strings.ContainsRune(second, c) {
			return c, true
		}
	}

	return 0, false
}

func FindSharedItemInGroup(first string, second string, third string) (rune, bool) {
	for _, c := range first {
		if strings.ContainsRune(second, c) && strings.ContainsRune(third, c) {
			return c, true
		}
	}

	return 0, false
}

func GetPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item - 'A' + 27)
	}

	return int(item - 'a' + 1)
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

func SumPriorities(lines []string) int {
	sum := 0

	for _, compartments := range lines {
		middle := len(compartments) / 2
		first := compartments[:middle]
		second := compartments[middle:]
		if shared, found := FindSharedItem(first, second); found {
			sum += GetPriority(shared)
		}
	}

	return sum
}

func SumGroupPriorities(lines []string) int {
	sum := 0

	for i := 0; i < len(lines)-2; i += 3 {
		a := lines[i]
		b := lines[i+1]
		c := lines[i+2]
		if shared, found := FindSharedItemInGroup(a, b, c); found {
			sum += GetPriority(shared)
		}
	}

	return sum
}

func main() {
	lines := ReadLines()

	// Part 1
	fmt.Println(SumPriorities(lines))

	// Part 2
	fmt.Println(SumGroupPriorities(lines))
}
