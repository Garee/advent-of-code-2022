package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	row     int
	col     int
	height  int
	visible bool
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

func ReadForestMatrix(lines []string) [][]Tree {
	trees := make([][]Tree, 0)
	for r, line := range lines {
		row := make([]Tree, 0)
		for col, c := range line {
			height, _ := strconv.Atoi(string(c))
			tree := Tree{
				row:     r,
				col:     col,
				height:  height,
				visible: false,
			}
			row = append(row, tree)
		}
		trees = append(trees, row)
	}

	return trees
}

func CountVisible(trees [][]Tree) (count int) {
	count += len(trees)*2 + len(trees[0])*2 - 4

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[i])-1 {
				trees[i][j].visible = true
			}
		}
	}

	for _, row := range trees {
		count += CountVisibleFacingLeft(row)
		count += CountVisibleFacingRight(row)
	}

	count += CountVisibleFacingDown(trees)
	count += CountVisibleFacingUp(trees)

	return count
}

func CountVisibleFacingDown(trees [][]Tree) (count int) {
	for i := 1; i < len(trees); i++ {
		for j := 1; j < len(trees[i]); j++ {
			if trees[i][j].height > trees[i-1][j-1].height && !trees[i][j].visible {
				fmt.Println(trees[i][j])
				trees[i][j].visible = true
				count++
			} else {
				return count
			}
		}
	}
	return count
}

func CountVisibleFacingUp(trees [][]Tree) (count int) {
	for i := len(trees) - 2; i > 0; i-- {
		for j := len(trees[0]) - 2; j > 0; j-- {
			if trees[i][j].height > trees[i+1][j+1].height && !trees[i][j].visible {
				fmt.Println(trees[i][j])
				trees[i][j].visible = true
				count++
			} else {
				return count
			}
		}
	}
	return count
}

func CountVisibleFacingRight(trees []Tree) (count int) {
	for i := 1; i < len(trees); i++ {
		if trees[i].height > trees[i-1].height && !trees[i].visible {
			fmt.Println(trees[i], "right")
			trees[i].visible = true
			count++
		} else {
			return count
		}
	}
	return count
}

func CountVisibleFacingLeft(trees []Tree) (count int) {
	for i := len(trees) - 2; i > 0; i-- {
		if trees[i].height > trees[i+1].height && !trees[i].visible {
			fmt.Println(trees[i], "left")
			trees[i].visible = true
			count++
		} else {
			return count
		}
	}
	return count
}

func main() {
	lines := ReadLines()
	forest := ReadForestMatrix(lines)

	// Part 1
	fmt.Println(CountVisible(forest))
}
