package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name string
	yell Yell
}

type Yell struct {
	lhs    *string
	rhs    *string
	op     *string
	result *int
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

func ParseMonkeys(lines []string) map[string]Monkey {
	monkeys := make(map[string]Monkey, 0)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		name := strings.ReplaceAll(tokens[0], ":", "")

		yell := Yell{}
		if len(tokens) > 2 {
			yell.lhs = &tokens[1]
			yell.op = &tokens[2]
			yell.rhs = &tokens[3]
		} else {
			n, _ := strconv.Atoi(tokens[1])
			yell.result = &n
		}

		monkey := Monkey{name, yell}
		monkeys[name] = monkey
	}
	return monkeys
}

func ComputeYellResult(monkey Monkey, monkeys map[string]Monkey) int {
	yell := monkey.yell
	if yell.result != nil {
		return *yell.result
	}

	lhs := ComputeYellResult(monkeys[*yell.lhs], monkeys)
	rhs := ComputeYellResult(monkeys[*yell.rhs], monkeys)

	switch *yell.op {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	default:
		return 0
	}
}

func main() {
	lines := ReadLines()
	monkeys := ParseMonkeys(lines)

	// Part 1
	result := ComputeYellResult(monkeys["root"], monkeys)
	fmt.Println(result)
}
