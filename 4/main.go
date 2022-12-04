package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ConvToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, _ := strconv.Atoi(s)
		ints[i] = n
	}
	return ints
}

func ContainsRange(range1 []int, range2 []int) bool {
	return (range1[0] <= range2[0] && range1[1] >= range2[1]) || (range2[0] <= range1[0] && range2[1] >= range1[1])
}

func ContainsOverlap(range1 []int, range2 []int) bool {
	return (range1[0] >= range2[0] && range1[0] <= range2[1]) || (range1[1] <= range2[1] && range1[1] >= range2[0]) || (range2[0] >= range1[0] && range2[0] <= range1[1]) || (range2[1] <= range1[1] && range2[1] >= range1[0])
}

func CountFullyContains(lines []string) (fully int, overlap int) {
	for _, line := range lines {
		elves := strings.Split(line, ",")
		elf1 := ConvToInts(strings.Split(elves[0], "-"))
		elf2 := ConvToInts(strings.Split(elves[1], "-"))
		if ContainsRange(elf1, elf2) {
			fully++
		}
		if ContainsOverlap(elf1, elf2) {
			overlap++
		}
	}

	return fully, overlap
}

func main() {
	lines := ReadLines()
	fully, overlap := CountFullyContains(lines)

	// Part 1
	fmt.Println(fully)

	// Part 2
	fmt.Println(overlap)
}
