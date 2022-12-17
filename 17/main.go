package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coord struct {
	x int
	y int
}

type Offset struct {
	x int
	y int
}

type Rock struct {
	kind   string
	coords []Coord
}
type Tower struct {
	width int
	rocks []Rock
}

var RockPattern = []string{"-", "+", "l", "i", "o"}

var RockOffsets = map[string][]Offset{
	"-": {
		Offset{2, 0},
		Offset{3, 0},
		Offset{4, 0},
		Offset{5, 0},
	},
	"+": {
		Offset{3, 0},
		Offset{2, 1},
		Offset{3, 1},
		Offset{4, 1},
		Offset{3, 2},
	},
	"l": {
		Offset{2, 0},
		Offset{3, 0},
		Offset{4, 0},
		Offset{4, 1},
		Offset{4, 2},
	},
	"i": {
		Offset{2, 0},
		Offset{2, 1},
		Offset{2, 2},
		Offset{2, 3},
	},
	"o": {
		Offset{2, 0},
		Offset{3, 0},
		Offset{2, 1},
		Offset{3, 1},
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

func Normalize(rocks []Rock) []Rock {
	min := math.MaxInt
	for _, rock := range rocks {
		for _, coord := range rock.coords {
			if coord.y < min {
				min = coord.y
			}
		}
	}

	for _, rock := range rocks {
		for _, coord := range rock.coords {
			coord.y -= min
		}
	}

	return rocks
}

func SimulateFallingRocks(rocks int, jets string) Tower {
	shapes := make([]Rock, 0)

	tower := Tower{
		width: Width,
		rocks: shapes,
	}

	jet := true
	rockIdx, jetIdx := 0, 0
	x, y := 0, BottomMargin
	prevRocks, prevHeight := 0, 0

	// While there are rocks remaining
	for rocks > 0 {
		// Add a new rock to the tower
		rockId := RockPattern[rockIdx]
		rockIdx = (rockIdx + 1) % len(RockPattern)

		coords := make([]Coord, 0)
		for _, offset := range RockOffsets[rockId] {
			coords = append(coords, Coord{
				x: x + offset.x,
				y: y + offset.y,
			})
		}

		rock := &Rock{rockId, coords}
		tower.rocks = append(tower.rocks, *rock)

		// While it is falling
		rockFall := true
		for rockFall {
			// If a gas jet occurs
			if jet {
				dir := string(jets[jetIdx])
				jetIdx = (jetIdx + 1) % len(jets)

				move := true
				if dir == ">" {

					for _, coord := range rock.coords {
						if coord.x+1 >= tower.width {
							move = false
							break
						}

						for j, other := range tower.rocks {
							if j == len(tower.rocks)-1 {
								continue
							}

							for _, otherCoord := range other.coords {
								if coord.x+1 == otherCoord.x && coord.y == otherCoord.y {
									move = false
									break
								}
							}
						}

						if !move {
							break
						}
					}

					if move {
						for i := range rock.coords {
							rock.coords[i].x += 1
						}
						rockFall = true
					}
				} else {
					for _, coord := range rock.coords {
						if coord.x-1 < 0 {
							move = false
							break
						}

						for j, other := range tower.rocks {
							if j == len(tower.rocks)-1 {
								continue
							}

							for _, otherCoord := range other.coords {
								if coord.x-1 == otherCoord.x && coord.y == otherCoord.y {
									move = false
									break
								}
							}
						}

						if !move {
							break
						}
					}

					if move {
						for i := range rock.coords {
							rock.coords[i].x -= 1
						}
						rockFall = true
					}
				}
			} else {
				for _, coord := range rock.coords {
					if coord.y-1 < 0 {
						rockFall = false
						break
					}

					for j, other := range tower.rocks {
						if j == len(tower.rocks)-1 {
							continue
						}

						for _, otherCoord := range other.coords {
							if coord.y-1 == otherCoord.y && coord.x == otherCoord.x {
								rockFall = false
								break
							}
						}

						if !rockFall {
							break
						}
					}

					if !rockFall {
						break
					}
				}

				if rockFall {
					for i := range rock.coords {
						rock.coords[i].y -= 1
					}
				}
			}

			jet = !jet
		}

		maxY := 0
		for _, rock := range tower.rocks {
			for _, coord := range rock.coords {
				if coord.y > maxY {
					maxY = coord.y
				}
			}
		}

		y = maxY + BottomMargin + 1

		rocks--

		//fmt.Println(rockIdx, jetIdx)
		//if rockIdx == 2 && jetIdx == 10 {
		if rockIdx == 2 && jetIdx == 23 {
			nRocks := len(tower.rocks)
			towerHeight := CalcTowerHeight(tower)

			fmt.Println("#", nRocks, "H:", towerHeight, "Diff #:", nRocks-prevRocks, "Diff H:", towerHeight-prevHeight, "(", rockIdx, jetIdx, ")")
			prevRocks = nRocks
			prevHeight = towerHeight
		}
	}

	return tower
}

func PrintTower(tower Tower) {
	for r := CalcTowerHeight(tower); r >= 0; r-- {
		for c := 0; c < tower.width; c++ {
			containsShape := false

			for _, shape := range tower.rocks {
				for _, coord := range shape.coords {
					if coord.x == c && coord.y == r {
						fmt.Print("#")
						containsShape = true
						break
					}
				}

				if containsShape {
					break
				}
			}

			if !containsShape {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func CalcTowerHeight(tower Tower) int {
	max := 0
	for _, rock := range tower.rocks {
		for _, coord := range rock.coords {
			if coord.y > max {
				max = coord.y
			}
		}
	}
	return max + 1
}

func main() {
	lines := ReadLines()
	jets := lines[0]

	// Part 1
	tower := SimulateFallingRocks(2022, jets)
	height := CalcTowerHeight(tower)
	fmt.Println(height)

	// Cycle every 1690 rocks and 2548 height
	//tower = SimulateFallingRocks(1000000000000, jets)
	// height = 5403
	// nRocks := 3580
	// nRocksToGo := 1000000000000 - nRocks
	// additionalHeight := (nRocksToGo / 1690) * 2548

	height = 3057
	nRocks := 2014
	nRocksToGo := 1000000000000 - nRocks
	additionalHeight := int(math.Floor(float64(nRocksToGo)/float64(35))) * 53
	fmt.Println(height + additionalHeight)
	fmt.Println((1000000000000-62)/(97-62)*(153-100) + 100)
	fmt.Println((1000000000000-200)/(1890-200)*(2855-307) + 307)
}
