package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type RopePart struct {
	id      string
	visited []Position
	x       int
	y       int
}

type Move struct {
	direction string
	amount    int
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

func ParseMoves(lines []string) []Move {
	moves := make([]Move, 0)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		amount, _ := strconv.Atoi(tokens[1])
		move := Move{
			direction: tokens[0],
			amount:    amount,
		}
		moves = append(moves, move)
	}
	return moves
}

func InitSim(startY int, n int) []RopePart {
	knots := make([]RopePart, 0)
	for i := 0; i < n; i++ {
		knot := RopePart{
			id:      fmt.Sprint(i),
			visited: make([]Position, 0),
			x:       startY,
			y:       startY,
		}
		knot.visited = append(knot.visited, Position{x: startY, y: startY})
		knots = append(knots, knot)
	}
	return knots
}

func FindDimension(moves []Move) (max int) {
	for _, move := range moves {
		if move.amount > max {
			max = move.amount
		}
	}
	return max
}

func CatchupTail(head RopePart, tail RopePart) RopePart {
	for head.y-tail.y > 1 {
		if head.x-tail.x > 0 {
			tail.x += 1
		} else if tail.x-head.x > 0 {
			tail.x -= 1
		}

		tail.y += 1
		tail = MarkVisited(tail)
	}
	for tail.y-head.y > 1 {
		if head.x-tail.x > 0 {
			tail.x += 1
		} else if tail.x-head.x > 0 {
			tail.x -= 1
		}

		tail.y -= 1
		tail = MarkVisited(tail)
	}
	for head.x-tail.x > 1 {
		if head.y-tail.y > 0 {
			tail.y += 1
		} else if tail.y-head.y > 0 {
			tail.y -= 1
		}

		tail.x += 1
		tail = MarkVisited(tail)
	}
	for tail.x-head.x > 1 {
		if head.y-tail.y > 0 {
			tail.y += 1
		} else if tail.y-head.y > 0 {
			tail.y -= 1
		}

		tail.x -= 1
		tail = MarkVisited(tail)
	}

	return tail
}

func MarkVisited(part RopePart) RopePart {
	for _, pos := range part.visited {
		if pos.x == part.x && pos.y == part.y {
			return part
		}
	}

	part.visited = append(part.visited, Position{x: part.x, y: part.y})
	return part
}
func PerformMove(move Move, head RopePart, tail RopePart) (RopePart, RopePart) {
	switch move.direction {
	case "R":
		for move.amount > 0 {
			head.x += 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
	case "L":
		for move.amount > 0 {
			head.x -= 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
	case "U":
		for move.amount > 0 {
			head.y -= 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
	case "D":
		for move.amount > 0 {
			head.y += 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
	default:
		break
	}

	return head, tail
}

func PerformMoves(moves []Move, head RopePart, tail RopePart) (RopePart, RopePart) {
	for _, move := range moves {
		head, tail = PerformMove(move, head, tail)
	}
	return head, tail
}

func PerformMoveMany(move Move, knots []RopePart, dim int) []RopePart {
	switch move.direction {
	case "R":
		for move.amount > 0 {
			knots[0].x += 1
			knots[0] = MarkVisited(knots[0])
			for i := 1; i < len(knots); i++ {
				knots[i] = CatchupTail(knots[i-1], knots[i])
			}
			move.amount--
		}
	case "L":
		for move.amount > 0 {
			knots[0].x -= 1
			knots[0] = MarkVisited(knots[0])
			for i := 1; i < len(knots); i++ {
				knots[i] = CatchupTail(knots[i-1], knots[i])
			}
			move.amount--
		}
	case "U":
		for move.amount > 0 {
			knots[0].y -= 1
			knots[0] = MarkVisited(knots[0])
			for i := 1; i < len(knots); i++ {
				knots[i] = CatchupTail(knots[i-1], knots[i])
			}
			move.amount--
		}
	case "D":
		for move.amount > 0 {
			knots[0].y += 1
			knots[0] = MarkVisited(knots[0])
			for i := 1; i < len(knots); i++ {
				knots[i] = CatchupTail(knots[i-1], knots[i])
			}
			move.amount--
		}
	default:
		break
	}
	return knots
}

func PerformMovesMany(knots []RopePart, moves []Move, dim int) []RopePart {
	for _, move := range moves {
		knots = PerformMoveMany(move, knots, dim)
	}
	return knots
}

func PrintVisited(tail RopePart, dim int) {
	for i := 0; i < dim*2; i++ {
		for j := 0; j < dim*2; j++ {
			found := false
			for _, pos := range tail.visited {
				if pos.x == j && pos.y == i {
					fmt.Print("#")
					found = true
				}
			}
			if !found {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func PrintState(knots []RopePart, dim int) {
	for i := 0; i < dim*2; i++ {
		for j := 0; j < dim*2; j++ {

			found := false
			for k, knot := range knots {
				if knot.x == j && knot.y == i {
					found = true
					fmt.Print(k)
					break
				}
			}
			if !found {
				if i == dim && j == dim {
					fmt.Print("s")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	lines := ReadLines()

	// Part 1
	moves := ParseMoves(lines)
	dim := FindDimension(moves)
	knots := InitSim(dim-1, 2)
	head, tail := knots[0], knots[1]
	_, tail = PerformMoves(moves, head, tail)
	fmt.Println(len(tail.visited))

	// Part 2
	moves = ParseMoves(lines)
	dim = FindDimension(moves)
	knots = InitSim(dim-1, 10)
	knots = PerformMovesMany(knots, moves, dim-1)
	tail = knots[len(knots)-1]
	fmt.Println(len(tail.visited))
}
