package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Player struct {
	pos    Coord
	target Coord
}

type Blizzard struct {
	pos Coord
	dir rune
}

type ValleyState struct {
	blizzards []Blizzard
	start     Coord
	end       Coord
	width     int
	height    int
}

func SimulateValleyStates(player Player, blizzards []Blizzard, valley []string, mins int) map[int]ValleyState {
	start, end := player.pos, player.target
	width, height := len(valley[0]), len(valley)
	statesByMin := make(map[int]ValleyState, 0)
	statesByMin[0] = ValleyState{blizzards, start, end, width, height}

	for m := 1; m < mins; m++ {
		blizzards = MoveBlizzards(blizzards, valley)
		statesByMin[m] = ValleyState{blizzards, start, end, width, height}
	}

	return statesByMin
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

func SurveyValley(valley []string) (Player, []Blizzard) {
	player := Player{}
	blizzards := make([]Blizzard, 0)
	for r, line := range valley {
		for c, ch := range line {
			if r == 0 && ch == '.' {
				player.pos = Coord{c, r}
			} else if r == len(valley)-1 && ch == '.' {
				player.target = Coord{c, r}
			} else if strings.Contains("<>^v", string(ch)) {
				blizzards = append(blizzards, Blizzard{Coord{c, r}, ch})
			}
		}
	}
	return player, blizzards
}

func ContainsBlizzard(x int, y int, blizzards []Blizzard) (Blizzard, bool) {
	for _, blizzard := range blizzards {
		if blizzard.pos.x == x && blizzard.pos.y == y {
			return blizzard, true
		}
	}
	return Blizzard{}, false
}

func MoveBlizzards(blizzards []Blizzard, valley []string) []Blizzard {
	result := make([]Blizzard, 0)
	for _, blizzard := range blizzards {
		switch blizzard.dir {
		case '>':
			blizzard.pos.x += 1
			if valley[blizzard.pos.y][blizzard.pos.x] == '#' {
				blizzard.pos.x = 1
			}
		case '<':
			blizzard.pos.x -= 1
			if valley[blizzard.pos.y][blizzard.pos.x] == '#' {
				blizzard.pos.x = len(valley[0]) - 2
			}
		case '^':
			blizzard.pos.y -= 1
			if valley[blizzard.pos.y][blizzard.pos.x] == '#' {
				if valley[len(valley)-1][blizzard.pos.x] == '.' {
					blizzard.pos.y = len(valley) - 1
				} else {
					blizzard.pos.y = len(valley) - 2
				}
			}
		case 'v':
			blizzard.pos.y += 1
			if valley[blizzard.pos.y][blizzard.pos.x] == '#' {
				if valley[0][blizzard.pos.x] == '.' {
					blizzard.pos.y = 0
				} else {
					blizzard.pos.y = 1
				}
			}
		}
		result = append(result, blizzard)
	}
	return result
}

func PathContains(path []Coord, p Coord) bool {
	for _, coord := range path {
		if coord.x == p.x && coord.y == p.y {
			return true
		}
	}
	return false
}

func ReachGoal(states map[int]ValleyState, mins int, start Coord, end Coord) int {
	path := []Coord{start}
	for !PathContains(path, end) && len(path) > 0 {
		mins++
		valley := states[mins]

		nextPath := make([]Coord, 0)
		for _, pos := range path {
			x, y := pos.x, pos.y

			steps := []Coord{
				{x, y - 1}, // up
				{x, y + 1}, // down
				{x - 1, y}, // left
				{x + 1, y}, // right
				{x, y},     // wait
			}

			for _, step := range steps {
				if step.x <= 0 || step.x >= valley.width-1 {
					continue
				}

				if step.y < 0 || step.y > valley.height-1 {
					continue
				}

				if step.y == 0 && start.y == 0 {
					if step.x != start.x {
						continue
					}
				}

				if step.y == 0 && end.y == 0 {
					if step.x != end.x {
						continue
					}
				}

				if step.y == valley.height-1 && start.y == valley.height-1 {
					if step.x != start.x {
						continue
					}
				}

				if step.y == valley.height-1 && end.y == valley.height-1 {
					if step.x != end.x {
						continue
					}
				}

				if _, hit := ContainsBlizzard(step.x, step.y, valley.blizzards); hit {
					continue
				}

				if !PathContains(nextPath, step) {
					nextPath = append(nextPath, step)
				}
			}
		}

		path = nextPath
	}

	return mins
}

func main() {
	valley := ReadLines()
	player, blizzards := SurveyValley(valley)
	states := SimulateValleyStates(player, blizzards, valley, 1024)

	// Part 1
	minsToEnd := ReachGoal(states, 0, player.pos, player.target)

	// Part 2
	t1 := ReachGoal(states, minsToEnd, player.target, player.pos)
	minsBackToStart := t1 - minsToEnd
	t2 := ReachGoal(states, t1, player.pos, player.target)
	minsBackToEnd := t2 - t1
	fmt.Println(minsToEnd + minsBackToStart + minsBackToEnd)
}
