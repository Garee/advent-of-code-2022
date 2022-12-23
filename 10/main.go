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
		case "noop":
			instructions = append(instructions, Inst{name: name, val: 0, cycle: 1})
		default:
			break
		}
	}
	return instructions
}

func SimulateSignal(instructions []Inst) []int {
	signal := make([]int, 0)

	x := 1
	for _, instr := range instructions {
		for instr.cycle > 0 {
			instr.cycle--
			signal = append(signal, x)
		}
		x += instr.val
	}

	return signal
}

func SimulateSignalStrengths(signal []int) []int {
	strengths := make([]int, 0)

	for cycle, s := range signal {
		strengths = append(strengths, s*(cycle+1))
	}

	return strengths
}

func Draw(signal []int) {
	width := 40
	height := 6

	pixels := make([]string, width*height)
	for i := range pixels {
		pixels[i] = "."
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			s := signal[row*width+col]
			if col == s || col == s-1 || col == s+1 {
				pixels[row*width+col] = "#"
			}
		}
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			pixel := pixels[row*width+col]
			fmt.Print(pixel)
		}
		fmt.Println()
	}
}

func main() {
	lines := ReadLines()
	instructions := ParseInstructions(lines)
	signal := SimulateSignal(instructions)
	strengths := SimulateSignalStrengths(signal)

	// Part 1
	sum := strengths[19] + strengths[59] + strengths[99] + strengths[139] + strengths[179] + strengths[219]
	fmt.Println(sum)

	// Part 2
	Draw(signal)
}
