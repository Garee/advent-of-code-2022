package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/gammazero/deque"
)

const Start = 'S'
const Best = 'E'

type Location struct {
	elevation int
	l         *Location
	r         *Location
	u         *Location
	d         *Location
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

func CreateLocations(lines []string) (locations []*Location, start *Location, end *Location) {
	height := len(lines)
	width := len(lines[0])
	locations = make([]*Location, width*height)
	for i := 0; i < len(locations); i++ {
		locations[i] = &Location{}
	}

	for row := 0; row < height; row++ {
		line := lines[row]
		for col := 0; col < width; col++ {
			location := locations[row*width+col]

			if line[col] == Start {
				location.elevation = int('a')
				start = location
			} else if line[col] == Best {
				location.elevation = int('z')
				end = location
			} else {
				location.elevation = int(line[col])
			}

			if col > 0 {
				l := locations[row*width+col-1]
				if l.elevation-location.elevation <= 1 {
					location.l = l
				}
			}

			if col < width-1 {
				r := locations[row*width+col+1]
				r.elevation = int(line[col+1])
				if r.elevation-location.elevation <= 1 {
					location.r = r
				}
			}

			if row > 0 {
				u := locations[(row-1)*width+col]
				if u.elevation-location.elevation <= 1 {
					location.u = u
				}
			}

			if row < height-1 {
				d := locations[(row+1)*width+col]

				elevation := int(lines[row+1][col])
				if elevation == Best {
					d.elevation = int('z')
				} else {
					d.elevation = elevation
				}

				if d.elevation-location.elevation <= 1 {
					location.d = d
				}
			}
		}
	}

	return locations, start, end
}

type V struct {
	pos  *Location
	dist int
}

func ShortestPathBfs(start *Location, end *Location) int {
	var q deque.Deque[V]
	q.PushBack(V{start, 0})

	seen := make(map[*Location]bool)
	for q.Len() > 0 {
		v := q.PopFront()
		pos, dist := v.pos, v.dist
		if pos == end {
			return dist
		}

		if seen[pos] {
			continue
		}
		seen[pos] = true

		adjacent := []*Location{
			pos.l, pos.u, pos.r, pos.d,
		}
		for _, adj := range adjacent {
			if adj != nil {
				q.PushBack(V{adj, dist + 1})
			}
		}
	}

	return math.MaxInt
}

func FindShortestRoute(locations []*Location, end *Location) int {
	shortestRoute := math.MaxInt

	for _, location := range locations {
		if location.elevation == int('a') {
			shortestPath := ShortestPathBfs(location, end)
			if shortestPath < shortestRoute {
				shortestRoute = shortestPath
			}
		}
	}

	return shortestRoute
}

func main() {
	lines := ReadLines()
	locations, start, end := CreateLocations(lines)

	// Part 1
	shortestPath := ShortestPathBfs(start, end)
	fmt.Println(shortestPath)

	// Part 2
	shortestRoute := FindShortestRoute(locations, end)
	fmt.Println(shortestRoute)
}
