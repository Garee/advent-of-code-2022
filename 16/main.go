package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Valve struct {
	name        string
	rate        int
	connections []*Valve
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

func ParseValvesAndTunnels(lines []string) ([]*Valve, *Valve) {
	valves := make([]*Valve, 0)
	lookup := make(map[string]*Valve, 0)

	for _, line := range lines {
		tokens := strings.Split(line, " ")

		name := tokens[1]

		rateStr := strings.ReplaceAll(tokens[4], ";", "")
		rateStr = strings.Split(rateStr, "=")[1]
		rate, _ := strconv.Atoi(rateStr)

		valve := Valve{
			name:        name,
			rate:        rate,
			connections: nil,
		}

		lookup[name] = &valve
		valves = append(valves, &valve)
	}

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		name := tokens[1]

		tokens = strings.Split(line, "valves ")
		if len(tokens) == 1 {
			tokens = strings.Split(line, "valve ")
		}

		connections := make([]*Valve, 0)
		for _, name := range strings.Split(tokens[1], ", ") {
			connections = append(connections, lookup[name])
		}

		lookup[name].connections = connections
	}

	return valves, lookup["AA"]
}

func CopyMap(m map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

func SimulatePressureRelease(curr *Valve, opened map[string]bool, min int, cache map[string]int) (pressure int) {
	key := curr.name + fmt.Sprint(opened) + fmt.Sprint(min)
	val, hit := cache[key]
	if hit {
		return val
	}

	if min <= 0 {
		return 0
	}

	for _, conn := range curr.connections {
		p := SimulatePressureRelease(conn, opened, min-1, cache)
		key := conn.name + fmt.Sprint(opened) + fmt.Sprint(min-1)
		cache[key] = p
		if p > pressure {
			pressure = p
		}
	}

	_, isOpen := opened[curr.name]
	if !isOpen && curr.rate > 0 && min > 0 {
		openedCopy := CopyMap(opened)
		openedCopy[curr.name] = true
		min--
		totalRate := min * curr.rate

		for _, conn := range curr.connections {
			p := SimulatePressureRelease(conn, openedCopy, min-1, cache)
			key := conn.name + fmt.Sprint(opened) + fmt.Sprint(min-1)
			cache[key] = p

			if totalRate+p > pressure {
				pressure = totalRate + p
			}
		}
	}

	return pressure
}

func main() {
	rand.Seed(time.Now().UnixNano())
	lines := ReadLines()

	// Part 1
	_, start := ParseValvesAndTunnels(lines)
	pressure := SimulatePressureRelease(start, make(map[string]bool, 0), 30, make(map[string]int))
	fmt.Println(pressure)
}
