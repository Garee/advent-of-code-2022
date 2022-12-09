package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	row         int
	col         int
	height      int
	visible     bool
	scenicScore int
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
				row:         r,
				col:         col,
				height:      height,
				visible:     false,
				scenicScore: 0,
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
	for col := 0; col < len(trees); col++ {
		max := trees[0][col].height
		for row := 0; row < len(trees[col]); row++ {
			if trees[row][col].height > max {
				max = trees[row][col].height

				if !trees[row][col].visible {
					trees[row][col].visible = true
					count++
				}
			}
		}
	}
	return count
}

func CountVisibleFacingUp(trees [][]Tree) (count int) {
	for col := 0; col < len(trees); col++ {
		max := trees[len(trees)-1][col].height

		for row := len(trees[col]) - 1; row >= 0; row-- {
			if trees[row][col].height > max {
				max = trees[row][col].height

				if !trees[row][col].visible {
					trees[row][col].visible = true
					count++
				}
			}
		}
	}
	return count
}

func CountVisibleFacingRight(trees []Tree) (count int) {
	max := trees[0].height
	for i := 0; i < len(trees); i++ {
		if trees[i].height > max {
			max = trees[i].height

			if !trees[i].visible {
				trees[i].visible = true
				count++
			}
		}
	}
	return count
}

func CountVisibleFacingLeft(trees []Tree) (count int) {
	max := trees[len(trees)-1].height
	for i := len(trees) - 1; i >= 0; i-- {
		if trees[i].height > max {
			max = trees[i].height

			if !trees[i].visible {
				trees[i].visible = true
				count++
			}
		}
	}
	return count
}

func CalculateScenicScores(trees [][]Tree) [][]Tree {
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			tree := trees[row][col]

			rightScore := 0
			i := 1
			if col < len(trees[row])-1 {
				rightScore++
				for col+i < len(trees)-1 && tree.height > trees[row][col+i].height {
					i++
					rightScore++
				}
			}

			leftScore := 0
			i = 1
			if col > 0 {
				leftScore++
				for col-i > 0 && tree.height > trees[row][col-i].height {
					i++
					leftScore++
				}
			}

			downScore := 0
			i = 1
			if row < len(trees)-1 {
				downScore++
				for row+i < len(trees)-1 && tree.height > trees[row+i][col].height {
					i++
					downScore++
				}
			}

			upScore := 0
			i = 1
			if row > 0 {
				upScore++
				for row-i > 0 && tree.height > trees[row-i][col].height {
					i++
					upScore++
				}
			}

			trees[row][col].scenicScore = rightScore * leftScore * downScore * upScore
		}
	}

	return trees
}

func FindMaxScenicScore(trees [][]Tree) int {
	max := 0
	for _, row := range trees {
		for _, tree := range row {
			if tree.scenicScore > max {
				max = tree.scenicScore
			}
		}
	}
	return max
}

func main() {
	lines := ReadLines()
	forest := ReadForestMatrix(lines)

	// Part 1
	fmt.Println(CountVisible(forest))

	// Part 2
	forest = CalculateScenicScores(forest)
	fmt.Println(FindMaxScenicScore(forest))
}
