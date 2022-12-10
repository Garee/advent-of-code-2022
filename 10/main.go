package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inst struct {
	name  string
	val   int
	cycle int
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

func ParseInstructions(lines []string) []Inst {
	instructions := make([]Inst, 0)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		name := tokens[0]
		switch name {
		case "addx":
			val, _ := strconv.Atoi(tokens[1])
			instructions = append(instructions, Inst{name: name, val: val, cycle: 2})
			break
		case "noop":
			instructions = append(instructions, Inst{name: name, val: 0, cycle: 1})
		default:
			break
		}
	}
	return instructions
}

func SimulateSignalStrengths(instructions []Inst) []int {
	strengths := make([]int, 0)

	x := 1
	cycle := 1
	for _, instr := range instructions {
		for instr.cycle > 0 {
			instr.cycle--
			strengths = append(strengths, x*cycle)
			cycle++
		}
		x += instr.val
	}

	return strengths
}

func main() {
	lines := ReadLines()
	instructions := ParseInstructions(lines)
	strengths := SimulateSignalStrengths(instructions)

	// Part 1
	sum := strengths[19] + strengths[59] + strengths[99] + strengths[139] + strengths[179] + strengths[219]
	fmt.Println(sum)
}
