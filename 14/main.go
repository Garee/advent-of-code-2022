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

const xOffset = 490

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

func CreateCave(paths []Path, width int, height int) ([]string, int) {
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

func DropSand(cave []string, floor int, from Pos) ([]string, int) {
	sand := from
	count := 0

	cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "O" + cave[sand.y][sand.x-xOffset+1:]

	for i := 0; i < 30; i++ {
		DrawCave(cave)

		if sand.y > floor {
			return cave, count
		}

		if cave[sand.y+1][sand.x-xOffset] == '#' || cave[sand.y+1][sand.x-xOffset] == 'O' {
			if sand.x-xOffset-1 < 0 {
				return cave, count
			} else if cave[sand.y+1][sand.x-xOffset-1] == '#' || cave[sand.y+1][sand.x-xOffset-1] == 'O' {
				if cave[sand.y+1][sand.x-xOffset+1] == '#' || cave[sand.y+1][sand.x-xOffset+1] == 'O' {
					sand = from
					count++
				} else {
					cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "." + cave[sand.y][sand.x-xOffset+1:]

					sand.y++
					sand.x = sand.x - 1

					cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "O" + cave[sand.y][sand.x-xOffset+1:]
				}
			} else {
				cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "." + cave[sand.y][sand.x-xOffset+1:]

				sand.y++
				sand.x = sand.x - 1

				cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "O" + cave[sand.y][sand.x-xOffset+1:]
			}
		} else {
			cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "." + cave[sand.y][sand.x-xOffset+1:]

			sand.y++

			cave[sand.y] = cave[sand.y][:sand.x-xOffset] + "O" + cave[sand.y][sand.x-xOffset+1:]
		}
	}

	return cave, count
}

func main() {
	lines := ReadLines()
	paths := ParsePaths(lines)

	// Part 1
	width, height := 20, 15
	cave, floor := CreateCave(paths, width, height)
	_, count := DropSand(cave, floor, Pos{500, 0})
	fmt.Println(count)
}
