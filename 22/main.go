package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLines() []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func FindStartPos(m []string) (int, int) {
	x, y := 0, 0

	for r := 0; r < len(m); r++ {
		x = strings.Index(m[r], ".")
		if x != -1 {
			y = r
			break
		}
	}

	return x, y
}

func ParseMap(lines []string) ([]string, string) {
	m := lines[:len(lines)-2]
	path := lines[len(lines)-1]
	return m, path
}

func GetRowBoundaries(m []string) (map[int]int, map[int]int) {
	left := make(map[int]int, 0)
	right := make(map[int]int, 0)
	for r := 0; r < len(m); r++ {
		ldot := strings.Index(m[r], ".")
		lhash := strings.Index(m[r], "#")
		if ldot != -1 && lhash == -1 {
			left[r] = ldot
		} else if ldot < lhash {
			left[r] = ldot
		} else {
			left[r] = -1
		}

		rdot := strings.LastIndex(m[r], ".")
		rhash := strings.LastIndex(m[r], "#")
		if rdot != -1 && rhash == -1 {
			right[r] = rdot
		} else if rdot > rhash {
			right[r] = rdot
		} else {
			right[r] = -1
		}
	}

	return left, right
}

func GetColBoundaries(m []string) (map[int]int, map[int]int) {
	up := make(map[int]int, 0)
	down := make(map[int]int, 0)

	for col := 0; col < len(m[0]); col++ {
		dot, hash := -1, -1
		for r := 0; r < len(m); r++ {
			if col >= len(m[r]) {
				continue
			}

			val := m[r][col]
			if val == '.' {
				dot = r
			} else if val == '#' {
				hash = r
			}
		}

		if dot != -1 && hash == -1 {
			down[col] = dot
		} else if dot > hash {
			down[col] = dot
		} else {
			down[col] = -1
		}

		dot, hash = -1, -1
		for r := len(m) - 1; r >= 0; r-- {
			if col >= len(m[r]) {
				continue
			}

			val := m[r][col]
			if val == '.' {
				dot = r
			} else if val == '#' {
				hash = r
			}
		}

		if dot != -1 && hash == -1 {
			up[col] = dot
		} else if dot < hash {
			up[col] = dot
		} else {
			up[col] = -1
		}
	}

	return up, down
}

func SimulatePath(m []string, path string, cube bool) (int, int, int) {
	x, y := FindStartPos(m)
	bl, br := GetRowBoundaries(m)
	bu, bd := GetColBoundaries(m)
	moveX, moveY := 1, 0

	num := ""
	for i, move := range path {
		_, err := strconv.Atoi(string(move))
		if err == nil {
			num += string(move)
			if i != len(path)-1 {
				continue
			}
		}
		n, _ := strconv.Atoi(num)
		num = ""
		for n > 0 {
			nextX := x + moveX
			nextY := y + moveY
			nmoveX := moveX
			nmoveY := moveY

			if !cube {
				if moveX < 0 {
					if nextX < 0 || nextX < bl[y] {
						if br[y] != -1 {
							nextX = br[y]
						} else {
							nextX = x
						}
					}
				} else if moveX > 0 {
					if nextX > br[y] && br[y] != -1 {
						if bl[y] != -1 {
							nextX = bl[y]
						} else {
							nextX = x
						}
					}
				}

				if moveY < 0 {
					if nextY < 0 || nextY < bu[x] {
						if bd[x] != -1 {
							nextY = bd[x]
						} else {
							nextY = y
						}
					}
				} else if moveY > 0 {
					if nextY > bd[x] && bd[x] != -1 {
						if bu[x] != -1 {
							nextY = bu[x]
						} else {
							nextY = y
						}
					}
				}
			} else {
				region1 := x >= 50 && x < 100 && y < 50
				region2 := x < 50 && y >= 100 && y < 150
				region3 := x < 50 && y >= 150
				region4 := x >= 50 && x < 100 && y >= 50 && y < 100
				region5 := x >= 100 && y < 50
				region6 := x >= 50 && x < 100 && y >= 100 && y < 150
				if moveX < 0 {
					if nextX < 0 || nextX < bl[y] {
						if region1 {
							// Region 1 to 2, se
							nextX = 0
							nextY = 149 - y
							nmoveX = 1
							nmoveY = 0
						} else if region4 {
							// Region 4 to 2, ss
							nextX = y - 50
							nextY = 100
							nmoveX = 0
							nmoveY = 1
						} else if region2 {
							// Region 2 to 1, se
							nextX = 50
							nextY = 149 - y
							nmoveX = 1
							nmoveY = 0
						} else if region3 {
							// Region 3 to 1, ss
							nextX = 50 + (y - 150)
							nextY = 0
							nmoveX = 0
							nmoveY = 1
						}
					}
				} else if moveX > 0 {
					if nextX > br[y] && br[y] != -1 {
						if region5 {
							// Region 5 to 6, ee
							nextX = 99
							nextY = 100 + (49 - y)
							nmoveX = -1
							nmoveY = 0
						} else if region4 {
							// Region 4 to 5, ss
							nextX = 100 + (y - 50)
							nextY = 49
							nmoveX = 0
							nmoveY = -1
						} else if region6 {
							// Region 6 to 5, ee
							nextX = 149
							nextY = 149 - y
							nmoveX = -1
							nmoveY = 0
						} else if region3 {
							// Region 3 to 6
							nextX = 50 + (y - 150)
							nextY = 149
							nmoveX = 0
							nmoveY = -1
						}
					}
				}

				if moveY < 0 {
					if nextY < 0 || nextY < bu[x] {
						if region1 {
							// Region 1 to 3
							nextX = 0
							nextY = 150 + (x - 50)
							nmoveX = 1
							nmoveY = 0
						} else if region5 {
							// Region 5 to 3
							nextX = x - 100
							nextY = 199
							nmoveX = 0
							nmoveY = -1
						} else if region2 {
							// Region 2 to 4
							nextX = 50
							nextY = 50 + x
							nmoveX = 1
							nmoveY = 0
						}
					}
				} else if moveY > 0 {
					if nextY > bd[x] && bd[x] != -1 {
						if region3 {
							// Region 3 to 5
							nextX = 100 + x
							nextY = 0
							nmoveX = 0
							nmoveY = 1
						} else if region6 {
							// Region 6 to 3
							nextX = 49
							nextY = 150 + (x - 50)
							nmoveX = -1
							nmoveY = 0
						} else if region5 {
							// Region 5 to 4
							nextX = 99
							nextY = 50 + (x - 100)
							nmoveX = -1
							nmoveY = 0
						}
					}
				}
			}

			if m[nextY][nextX] == '.' {
				x = nextX
				y = nextY
				moveX = nmoveX
				moveY = nmoveY
			}

			n--
		}

		switch move {
		case 'R':
			if moveX == 1 && moveY == 0 {
				moveX = 0
				moveY = 1
			} else if moveX == 0 && moveY == 1 {
				moveX = -1
				moveY = 0
			} else if moveX == -1 && moveY == 0 {
				moveX = 0
				moveY = -1
			} else {
				moveX = 1
				moveY = 0
			}
		case 'L':
			if moveX == 1 && moveY == 0 {
				moveX = 0
				moveY = -1
			} else if moveX == 0 && moveY == -1 {
				moveX = -1
				moveY = 0
			} else if moveX == -1 && moveY == 0 {
				moveX = 0
				moveY = 1
			} else {
				moveX = 1
				moveY = 0
			}
		default:
			// Do nothing
		}
	}

	facing := 0
	if moveX == 0 && moveY == 1 {
		facing = 1
	} else if moveX == -1 && moveY == 0 {
		facing = 2
	} else if moveX == 0 && moveY == -1 {
		facing = 3
	}

	return y + 1, x + 1, facing
}

func FinalPassword(row int, col int, facing int) int {
	return row*1000 + col*4 + facing
}

func main() {
	lines := ReadLines()
	m, path := ParseMap(lines)

	// Part 1
	row, col, facing := SimulatePath(m, path, false)
	password := FinalPassword(row, col, facing)
	fmt.Println(password)

	// Part 2
	row, col, facing = SimulatePath(m, path, true)
	password = FinalPassword(row, col, facing)
	fmt.Println(password)
}
