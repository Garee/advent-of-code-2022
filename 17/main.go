package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

var Pattern = []string{"-", "+", "l", "i", "o"}

var Shapes = map[string][]string{
	"-": {
		".......",
		".......",
		".......",
		"..####.",
	},
	"+": {
		".......",
		".......",
		".......",
		"...#...",
		"..###..",
		"...#...",
	},
	"l": {
		".......",
		".......",
		".......",
		"....#..",
		"....#..",
		"..###..",
	},
	"i": {
		".......",
		".......",
		".......",
		"..#....",
		"..#....",
		"..#....",
		"..#....",
	},
	"o": {
		".......",
		".......",
		".......",
		"..##...",
		"..##...",
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

func SimulateFallingRocks(rocks int, jets string) ([]string, int) {
	tower := make([]string, 0)

	rockFall := true
	jet := true
	rockIdx, jetIdx := 0, 0

	for rocks > 0 {
		rock := Shapes[Pattern[rockIdx]]
		tower = append(tower, rock...)
		rocks--

		rockIdx++
		if rockIdx >= len(Pattern) {
			rockIdx = 0
		}

		for rockFall {
			if jet {
				// TODO:

				jetIdx++
				if jetIdx >= len(jets) {
					jetIdx = 0
				}
				jet = false
			} else {
				height := 0
				top := rock[0]
				for i, line := range rock {
					if strings.Contains(line, "#") {
						height = len(rock) - i + 1
						top = line
					}
				}
			}
		}
	}

	return tower, len(tower)
}

func main() {
	lines := ReadLines()
	jets := lines[0]
	tower, height := SimulateFallingRocks(2, jets)

	for r := 0; r < len(tower); r++ {
		fmt.Println(tower[r])
	}

	fmt.Println(height)
}
