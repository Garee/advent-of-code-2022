package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Path []Pos

func ReadLines() []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func ParsePaths(lines []string) []Path {
	paths := make([]Path, 0)
	for _, line := range lines {
		path := make([]Pos, 0)
		tokens := strings.Split(line, " -> ")
		for _, token := range tokens {
			coord := strings.Split(token, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			path = append(path, Pos{x, y})
		}
		paths = append(paths, path)
	}
	return paths
}

func CreateCave(paths []Path, width int, height int, xOffset int) ([]string, int) {
	cave := make([]string, 0, height)

	for i := 0; i < height; i++ {
		row := ""
		for j := 0; j < width; j++ {
			row += "."
		}
		cave = append(cave, row)
	}

	floor := 0

	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			from := path[i-1]
			to := path[i]

			var start, end int
			if from.x != to.x {
				if from.x > to.x {
					start, end = to.x, from.x
				} else {
					start, end = from.x, to.x
				}

				for x := start; x <= end; x++ {
					y := to.y
					cave[y] = cave[y][:x-xOffset] + "#" + cave[y][x-xOffset+1:]
				}
			} else if from.y != to.y {
				if from.y > to.y {
					start, end = to.y, from.y
				} else {
					start, end = from.y, to.y
				}

				for y := start; y <= end; y++ {
					cave[y] = cave[y][:to.x-xOffset] + "#" + cave[y][to.x-xOffset+1:]
				}
			}

			if to.y > floor {
				floor = to.y
			}
		}
	}

	return cave, floor
}

func DrawCave(cave []string) {
	for _, row := range cave {
		fmt.Println(row)
	}
}

func DropSand(cave []string, floor int, from Pos, xOffset int, solidFloor bool) ([]string, int) {
	sand := from
	count := 0

	cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "O" + cave[sand.y][sand.x-xOffset+1:]

	for {
		if sand.y > floor && !solidFloor {
			return cave, count
		}

		x := sand.x - xOffset
		xl := x - 1
		xr := x + 1
		y := sand.y + 1

		hitFloor := solidFloor && y == floor+2

		// Check down.
		if cave[y][x] == '#' || cave[y][x] == 'O' || hitFloor {
			// Down is blocked.

			// Check left.
			if xl < 0 || cave[y][xl] == '#' || cave[y][xl] == 'O' || hitFloor {
				// Left is blocked.

				// Check right.
				if cave[y][xr] == '#' || cave[y][xr] == 'O' || hitFloor {
					// Right is blocked.
					if sand.x == from.x && sand.y == from.y {
						return cave, count + 1
					}

					sand = from
					count++
				} else {
					// Right is not blocked.
					sand.y = y
					sand.x += 1

					cave[y-1] = cave[y-1][:x] + "." + cave[y-1][x+1:]
					cave[y] = cave[y][:xr] + "O" + cave[y][xr+1:]
				}
			} else {
				// Left is not blocked
				sand.y = y
				sand.x -= 1

				cave[y-1] = cave[y-1][:x] + "." + cave[y-1][x+1:]
				cave[y] = cave[y][:xl] + "O" + cave[y][xl+1:]
			}
		} else {
			// Down is not blocked.
			sand.y = y
			cave[y-1] = cave[sand.y-1][:x] + "." + cave[y-1][x+1:]
			cave[y] = cave[y][:x] + "O" + cave[y][x+1:]
		}
	}
}

func main() {
	lines := ReadLines()
	paths := ParsePaths(lines)

	// Part 1
	width, height, xOffset := 1000, 1000, 0
	cave, floor := CreateCave(paths, width, height, xOffset)
	_, count := DropSand(cave, floor, Pos{500, 0}, xOffset, false)
	fmt.Println(count)

	// Part 2
	width, height, xOffset = 1000, 1000, 0
	cave, floor = CreateCave(paths, width, height, xOffset)
	_, count = DropSand(cave, floor, Pos{500, 0}, xOffset, true)
	fmt.Println(count)
}
