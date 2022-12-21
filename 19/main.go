package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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

func GetNextRobotToBuildInner(bp Blueprint, resources map[string]int, options []string, deps map[string][]string) (string, Cost) {
	for _, option := range options {
		cost := GetCost(option, bp)
		if Affordable(cost, resources) {
			return option, cost
		}

		dep, depCost := GetNextRobotToBuildInner(bp, resources, deps[option], deps)
		if dep != "" {
			return dep, depCost
		}
	}

	return "", nil
}

func GetNextRobotToBuild(bp Blueprint, resources map[string]int) (string, Cost) {
	options := []string{"geode", "obsidian", "clay", "ore"}

	deps := map[string][]string{
		"geode":    {"obsidian", "ore"},
		"obsidian": {"clay", "ore"},
		"clay":     {"ore"},
		"ore":      {},
	}

	return GetNextRobotToBuildInner(bp, resources, options, deps)
}

func CountOpenedGeodes(bp Blueprint, mins int, robots map[string]int, resources map[string]int) int {
	if mins <= 0 {
		return resources["geode"]
	}

	rob := CopyMap(robots)
	res := CopyMap(resources)

	robotToBuild, cost := GetNextRobotToBuild(bp, res)
	rob[robotToBuild] += 1
	for k := range cost {
		res[k] -= cost[k]
	}

	countWithRobot := 0
	if robotToBuild != "" {
		for k, v := range robots {
			res[k] += v
		}
		countWithRobot = CountOpenedGeodes(bp, mins-1, rob, res)
	}

	for k, v := range robots {
		resources[k] += v
	}

	countWithoutRobot := CountOpenedGeodes(bp, mins-1, robots, resources)
	return Max(countWithRobot, countWithoutRobot)
}

func QualityLevel(bp Blueprint, mins int) int {
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
	count := CountOpenedGeodes(bp, mins, robots, resources)
	return bp.id * count
}

func SumQualityLevels(bps []Blueprint, mins int) (sum int) {
	for _, bp := range bps {
		sum += QualityLevel(bp, mins)
	}
	return sum
}

func main() {
	lines := ReadLines()
	blueprints := ParseBlueprints(lines)

	// Part 1
	mins := 24
	fmt.Println(QualityLevel(blueprints[0], mins))
	fmt.Println(QualityLevel(blueprints[1], mins))
}
