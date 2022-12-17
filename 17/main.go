package main

import (
	"bufio"
	"fmt"
	"os"
)

var Pattern = []string{"-", "+", "l", "i", "o"}

var Shapes = map[string][]string{
	"-": {
		"..####.",
		".......",
		".......",
		".......",
	},
	"+": {
		"...#...",
		"..###..",
		"...#...",
		".......",
		".......",
		".......",
	},
	"l": {
		"....#..",
		"....#..",
		"..###..",
		".......",
		".......",
		".......",
	},
	"i": {
		"..#....",
		"..#....",
		"..#....",
		"..#....",
		".......",
		".......",
		".......",
	},
	"o": {
		"..##...",
		"..##...",
		".......",
		".......",
		".......",
	},
}

const Width = 7
const LeftMargin = 2
const BottomMargin = 3

func ReadLines() []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func SimulateFallingRocks(rocks int, jets string) int {
	tower := make([]string, 0)

	rockFall := true
	rockIdx, jetIdx := 0, 0

	for rocks > 0 {
		if rockFall {
			rock := Shapes[Pattern[rockIdx]]
			tower = append(tower, rock...)
			rocks--

			rockIdx++
			if rockIdx >= len(Pattern) {
				rockIdx = 0
			}
			rockFall = false
		} else {
			jetIdx++
			if jetIdx >= len(jets) {
				jetIdx = 0
			}
			rockFall = true
		}
	}

	for r := 0; r < len(tower); r++ {
		fmt.Println(tower[r])
	}

	return len(tower)
}

func main() {
	lines := ReadLines()
	jets := lines[0]
	height := SimulateFallingRocks(2022, jets)
	fmt.Println(height)
}
