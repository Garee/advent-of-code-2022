package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Elf struct {
	x int
	y int
}

type Coord struct {
	x int
	y int
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

func InitDirs() []string {
	return []string{"N", "S", "W", "E"}
}

func NextDirs(dirs []string) []string {
	return append(dirs[1:], dirs[0])
}

func ParseElves(lines []string) ([]*Elf, []string) {
	elves := make([]*Elf, 0)
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				elf := &Elf{x, y}
				elves = append(elves, elf)
			}
		}
	}
	return elves, lines
}

func GetElfLocations(elves []*Elf) map[Coord]*Elf {
	locs := make(map[Coord]*Elf, 0)
	for _, elf := range elves {
		coord := Coord{elf.x, elf.y}
		locs[coord] = elf
	}
	return locs
}

func HasNorthNeighbour(elf *Elf, currLocs map[Coord]*Elf) bool {
	coords := []Coord{
		{elf.x, elf.y - 1},     // N
		{elf.x - 1, elf.y - 1}, // NW
		{elf.x + 1, elf.y - 1}, // NE
	}
	for _, coord := range coords {
		if _, found := currLocs[coord]; found {
			return true
		}
	}
	return false
}

func HasSouthNeighbour(elf *Elf, currLocs map[Coord]*Elf) bool {
	coords := []Coord{
		{elf.x, elf.y + 1},     // S
		{elf.x - 1, elf.y + 1}, // SW
		{elf.x + 1, elf.y + 1}, // SE
	}
	for _, coord := range coords {
		if _, found := currLocs[coord]; found {
			return true
		}
	}
	return false
}

func HasWestNeighbour(elf *Elf, currLocs map[Coord]*Elf) bool {
	coords := []Coord{
		{elf.x - 1, elf.y - 1}, // NW
		{elf.x - 1, elf.y + 1}, // SW
		{elf.x - 1, elf.y},     // W
	}
	for _, coord := range coords {
		if _, found := currLocs[coord]; found {
			return true
		}
	}
	return false
}

func HasEastNeighbour(elf *Elf, currLocs map[Coord]*Elf) bool {
	coords := []Coord{
		{elf.x + 1, elf.y - 1}, // NE
		{elf.x + 1, elf.y + 1}, // SE
		{elf.x + 1, elf.y},     // E
	}
	for _, coord := range coords {
		if _, found := currLocs[coord]; found {
			return true
		}
	}
	return false
}

func HasNeighbour(elf *Elf, currLocs map[Coord]*Elf) bool {
	return HasNorthNeighbour(elf, currLocs) || HasSouthNeighbour(elf, currLocs) || HasWestNeighbour(elf, currLocs) || HasEastNeighbour(elf, currLocs)
}

func Simulate(elves []*Elf, dirs []string, rounds int) []*Elf {
	for rounds > 0 {
		currLocs := GetElfLocations(elves)
		nextLocs := make(map[*Elf]Coord, 0)

		for _, elf := range elves {
			if !HasNeighbour(elf, currLocs) {
				continue
			}

			for _, dir := range dirs {
				if _, found := nextLocs[elf]; found {
					break
				}

				switch dir {
				case "N":
					if !HasNorthNeighbour(elf, currLocs) {
						north := Coord{elf.x, elf.y - 1}
						nextLocs[elf] = north
					}
				case "S":
					if !HasSouthNeighbour(elf, currLocs) {
						south := Coord{elf.x, elf.y + 1}
						nextLocs[elf] = south
					}
				case "W":
					if !HasWestNeighbour(elf, currLocs) {
						west := Coord{elf.x - 1, elf.y}
						nextLocs[elf] = west
					}
				case "E":
					if !HasEastNeighbour(elf, currLocs) {
						east := Coord{elf.x + 1, elf.y}
						nextLocs[elf] = east
					}
				}
			}
		}

		for elf, coord := range nextLocs {
			canMove := true

			for other, otherCoord := range nextLocs {
				if elf == other {
					continue
				}

				if coord.x == otherCoord.x && coord.y == otherCoord.y {
					canMove = false
					break
				}
			}

			if canMove {
				elf.x = coord.x
				elf.y = coord.y
			}
		}

		dirs = NextDirs(dirs)
		rounds--
	}

	return elves
}

func GetBoundary(elves []*Elf) (int, int, int, int) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for _, elf := range elves {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}
	return minX, minY, maxX, maxY
}

func CountEmpty(elves []*Elf, draw bool) (count int) {
	minX, minY, maxX, maxY := GetBoundary(elves)
	for r := minY; r < maxY+1; r++ {
		for c := minX; c < maxX+1; c++ {
			found := false
			for _, elf := range elves {
				if elf.x == c && elf.y == r {
					found = true
					break
				}
			}

			if found {
				if draw {
					fmt.Print("#")
				}

			} else {
				count++
				if draw {
					fmt.Print(".")
				}

			}
		}
		if draw {
			fmt.Println()
		}

	}

	if draw {
		fmt.Println()
	}

	return count
}

func main() {
	lines := ReadLines()

	// Part 1
	elves, _ := ParseElves(lines)
	dirs := InitDirs()
	rounds := 10
	elves = Simulate(elves, dirs, rounds)
	count := CountEmpty(elves, false)
	fmt.Println(count)
}
