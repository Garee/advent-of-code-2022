package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord3D struct {
	x int
	y int
	z int
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

func ParseCoords(lines []string) []Coord3D {
	coords := make([]Coord3D, 0)
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])
		coords = append(coords, Coord3D{x, y, z})
	}
	return coords
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func CalcSurfaceArea(coords []Coord3D) (area int) {
	for _, coord := range coords {
		sides := 6
		for _, other := range coords {
			if other.x == coord.x && other.y == coord.y && other.z == coord.z {
				continue
			}

			xd := other.x - coord.x
			yd := other.y - coord.y
			zd := other.z - coord.z

			if Abs(xd) == 1 && yd == 0 && zd == 0 {
				sides--
			}

			if xd == 0 && Abs(yd) == 1 && zd == 0 {
				sides--
			}

			if xd == 0 && yd == 0 && Abs(zd) == 1 {
				sides--
			}
		}
		area += sides
	}
	return area
}

func main() {
	lines := ReadLines()
	coords := ParseCoords(lines)

	// Part 1
	area := CalcSurfaceArea(coords)
	fmt.Println(area)
}
