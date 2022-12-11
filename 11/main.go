package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	kind string
	rhs  string
}

type Test struct {
	divisibleBy int
	monkeyTrue  int
	monkeyFalse int
}

type Monkey struct {
	name      string
	items     []int
	operation Operation
	test      Test
	inspected int
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

func StrconvArrayToI(arr []string) []int {
	result := make([]int, 0)
	for _, s := range arr {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func ParseMonkeys(lines []string) []Monkey {
	monkeys := make([]Monkey, 0)
	for i := 0; i < len(lines); i += 7 {
		name := strings.ReplaceAll(lines[i], ":", "")
		items := strings.Split(strings.Split(lines[i+1], ": ")[1], ", ")

		tokens := strings.Split(lines[i+2], " ")
		kind := tokens[len(tokens)-2]
		rhs := tokens[len(tokens)-1]

		tokens = strings.Split(lines[i+3], " ")
		divisibleBy, _ := strconv.Atoi(tokens[len(tokens)-1])

		tokens = strings.Split(lines[i+4], " ")
		monkeyTrue, _ := strconv.Atoi(tokens[len(tokens)-1])

		tokens = strings.Split(lines[i+5], " ")
		monkeyFalse, _ := strconv.Atoi(tokens[len(tokens)-1])

		monkey := Monkey{
			name:  name,
			items: StrconvArrayToI(items),
			operation: Operation{
				kind: kind,
				rhs:  rhs,
			},
			test: Test{
				divisibleBy: divisibleBy,
				monkeyTrue:  monkeyTrue,
				monkeyFalse: monkeyFalse,
			},
			inspected: 0,
		}
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func SimulateMonkeyBusiness(monkeys []Monkey, rounds int) []Monkey {
	for i := 0; i < rounds; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.items {
				op := monkey.operation
				test := monkey.test

				rhs, err := strconv.Atoi(op.rhs)
				if err != nil {
					rhs = item
				}

				worry := item
				if op.kind == "*" {
					worry *= rhs
				} else if op.kind == "+" {
					worry += rhs
				}
				worry /= 3

				if worry%test.divisibleBy == 0 {
					monkeys[test.monkeyTrue].items = append(monkeys[test.monkeyTrue].items, worry)
				} else {
					monkeys[test.monkeyFalse].items = append(monkeys[test.monkeyFalse].items, worry)
				}

				monkeys[m].inspected++
			}

			monkeys[m].items = make([]int, 0)
		}
	}
	return monkeys
}

func main() {
	lines := ReadLines()
	monkeys := ParseMonkeys(lines)
	monkeys = SimulateMonkeyBusiness(monkeys, 20)

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	// Part 1
	fmt.Println(monkeys[0].inspected * monkeys[1].inspected)
}
