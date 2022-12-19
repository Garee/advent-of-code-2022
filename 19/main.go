package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var MaxGeodes = 0

type Cost = map[string]int

type Blueprint struct {
	id       int
	ore      Cost
	clay     Cost
	obsidian Cost
	geode    Cost
}

func GetCost(robot string, bp Blueprint) Cost {
	switch robot {
	case "ore":
		return bp.ore
	case "clay":
		return bp.clay
	case "obsidian":
		return bp.obsidian
	case "geode":
		return bp.geode
	default:
		return Cost{}
	}
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

func StrToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func ParseBlueprints(lines []string) []Blueprint {
	blueprints := make([]Blueprint, 0)
	for i, line := range lines {
		re := regexp.MustCompile(`\d+`)
		amounts := re.FindAllString(line, -1)

		blueprint := Blueprint{
			id:       i + 1,
			ore:      Cost{"ore": StrToInt(amounts[1])},
			clay:     Cost{"ore": StrToInt(amounts[2])},
			obsidian: Cost{"ore": StrToInt(amounts[3]), "clay": StrToInt(amounts[4])},
			geode:    Cost{"ore": StrToInt(amounts[5]), "obsidian": StrToInt(amounts[6])},
		}
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CopyMap(src map[string]int) map[string]int {
	m := make(map[string]int, 0)
	for k, v := range src {
		m[k] = v
	}
	return m
}

func Affordable(cost map[string]int, resources map[string]int) bool {
	for k, v := range cost {
		if resources[k]-v < 0 {
			return false
		}
	}
	return true
}

func MaxRequired(bp Blueprint, robot string) (max int) {
	if robot == "geode" {
		return math.MaxInt
	}
	if bp.geode[robot] > max {
		max = bp.geode[robot]
	}
	if bp.obsidian[robot] > max {
		max = bp.obsidian[robot]
	}
	if bp.clay[robot] > max {
		max = bp.clay[robot]
	}
	return max
}

func Fadd(geodes int, mins int) int {
	result := 0
	for i := 1; i < mins+1; i++ {
		result += geodes + i
	}
	return result
}

func CountOpenedGeodes(bp Blueprint, mins int, robots map[string]int, resources map[string]int, cache map[string]int) int {
	MaxGeodes = Max(MaxGeodes, resources["geode"])
	if mins <= 0 {
		return resources["geode"]
	}

	if resources["geode"]+Fadd(robots["geode"], mins) <= MaxGeodes {
		return resources["geode"]
	}

	key := fmt.Sprintf("%d_%d_%s_%s", bp.id, mins, robots, resources)
	if v, hit := cache[key]; hit {
		return v
	}

	maxCountWithRobot := 0
	options := []string{"geode", "obsidian", "clay", "ore"}
	for _, option := range options {
		maxRequired := MaxRequired(bp, option)
		if robots[option] >= maxRequired {
			continue
		}

		cost := GetCost(option, bp)
		if !Affordable(cost, resources) {
			continue
		}

		rob := CopyMap(robots)
		res := CopyMap(resources)

		rob[option] += 1
		for k := range cost {
			res[k] -= cost[k]
		}

		for k, v := range robots {
			res[k] += v
		}

		count := 0
		key := fmt.Sprintf("%d_%d_%s_%s", bp.id, mins-1, rob, res)
		if v, hit := cache[key]; hit {
			count = v
		} else {
			count = CountOpenedGeodes(bp, mins-1, rob, res, cache)
			cache[key] = count
		}

		maxCountWithRobot = Max(maxCountWithRobot, count)

		if option == "geode" {
			break
		}
	}

	for k, v := range robots {
		resources[k] += v
	}

	countWithoutRobot := 0
	key = fmt.Sprintf("%d_%d_%s_%s", bp.id, mins-1, robots, resources)
	if v, hit := cache[key]; hit {
		countWithoutRobot = v
	} else {
		countWithoutRobot = CountOpenedGeodes(bp, mins-1, robots, resources, cache)
		cache[key] = countWithoutRobot
	}

	return Max(maxCountWithRobot, countWithoutRobot)
}

func QualityLevel(bp Blueprint, mins int, cache map[string]int) (int, int) {
	robots := map[string]int{
		"ore":      1,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	resources := map[string]int{
		"ore":      0,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	MaxGeodes = 0
	count := CountOpenedGeodes(bp, mins, robots, resources, cache)
	return bp.id * count, count
}

func SumQualityLevels(bps []Blueprint, mins int, cache map[string]int) (sum int) {
	for _, bp := range bps {
		ql, _ := QualityLevel(bp, mins, cache)
		sum += ql
	}
	return sum
}

func main() {
	rand.Seed(time.Now().UnixNano())

	lines := ReadLines()
	blueprints := ParseBlueprints(lines)

	// Part 1
	cache := make(map[string]int, 0)
	fmt.Println(SumQualityLevels(blueprints, 24, cache))

	// Part 2
	_, first := QualityLevel(blueprints[0], 32, cache)
	_, second := QualityLevel(blueprints[1], 32, cache)
	_, third := QualityLevel(blueprints[2], 32, cache)
	fmt.Println(first * second * third)
}
