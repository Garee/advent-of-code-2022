package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Sensor struct {
	x      int
	y      int
	beacon Beacon
	dist   int
}

type Beacon struct {
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

func ParseSensorsAndBeacons(lines []string) ([]Sensor, []Beacon, int, int, int, int) {
	sensors := make([]Sensor, 0)
	beacons := make([]Beacon, 0)
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for _, line := range lines {
		re := regexp.MustCompile(`[-?\d*]+`)
		nums := re.FindAllString(line, -1)

		x, _ := strconv.Atoi(nums[2])
		y, _ := strconv.Atoi(nums[3])
		beacon := Beacon{x, y}
		beacons = append(beacons, beacon)
		minX = Min(minX, x)
		minY = Min(minY, y)
		maxX = Max(maxX, x)
		maxY = Max(maxY, y)

		x, _ = strconv.Atoi(nums[0])
		y, _ = strconv.Atoi(nums[1])
		sensor := Sensor{x, y, beacon, 0}
		sensors = append(sensors, sensor)
		minX = Min(minX, x)
		minY = Min(minY, y)
		maxX = Max(maxX, x)
		maxY = Max(maxY, y)
	}
	return sensors, beacons, minX, minY, maxX, maxY
}

func Draw(sensors []Sensor, beacons []Beacon, minX int, maxX int, r int) int {
	count := 0
	for c := minX * 10; c < maxX*10; c++ {
		blocked := false
		inRange := false

		for _, beacon := range beacons {
			if beacon.x == c && beacon.y == r {
				blocked = true
				break
			}
		}

		for _, sensor := range sensors {
			if sensor.x == c && sensor.y == r {
				blocked = true
				break
			}

			dist := ManDist(c, sensor.x, r, sensor.y)
			if dist <= sensor.dist {
				inRange = true
				break
			}
		}

		if inRange && !blocked {
			count++
		}
	}

	return count
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func CalcDistances(sensors []Sensor) []Sensor {
	for i := 0; i < len(sensors); i++ {
		sensors[i].dist = Dist(sensors[i], sensors[i].beacon)
	}
	return sensors
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ManDist(x1 int, x2 int, y1 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func Dist(a Sensor, b Beacon) int {
	return ManDist(a.x, b.x, a.y, b.y)
}

func main() {
	lines := ReadLines()
	sensors, beacons, minX, _, maxX, _ := ParseSensorsAndBeacons(lines)
	sensors = CalcDistances(sensors)
	fmt.Println(Draw(sensors, beacons, minX, maxX, 2000000))
}
