package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Sum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func CountCaloriesByElf() []int {
	calories_by_elf := []int{0}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calories_by_elf = append(calories_by_elf, 0)
		} else {
			calories, _ := strconv.Atoi(line)
			calories_by_elf[len(calories_by_elf)-1] += calories
		}
	}

	return calories_by_elf
}

func main() {
	// Part 1
	calories_by_elf := CountCaloriesByElf()
	sort.Ints(calories_by_elf)
	max := calories_by_elf[len(calories_by_elf)-1]
	fmt.Println(max)

	// Part 2
	total := Sum(calories_by_elf[len(calories_by_elf)-3:])
	fmt.Println(total)
}
