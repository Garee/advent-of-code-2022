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
			x:       0,
			y:       startY,
		}
		knot.visited = append(knot.visited, Position{x: 0, y: startY})
		knots = append(knots, knot)
	}
	return knots
}

func FindDimension(moves []Move) (max int) {
	for _, move := range moves {
		if move.direction == "U" || move.direction == "D" && move.amount > max {
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
		break
	case "L":
		for move.amount > 0 {
			head.x -= 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
		break
	case "U":
		for move.amount > 0 {
			head.y -= 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
		break
	case "D":
		for move.amount > 0 {
			head.y += 1
			head = MarkVisited(head)
			tail = CatchupTail(head, tail)
			move.amount--
		}
		break
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

func PerformMovesMany(knots []RopePart, moves []Move) []RopePart {
	for _, move := range moves {
		a, b := PerformMove(move, knots[0], knots[1])
		knots[0] = a
		knots[1] = b
		for i := 2; i < len(knots); i++ {
			knots[i] = CatchupTail(knots[i-1], knots[i])
		}
	}
	return knots
}

func main() {
	lines := ReadLines()
	moves := ParseMoves(lines)
	dim := FindDimension(moves)
	knots := InitSim(dim-1, 2)
	head, tail := knots[0], knots[1]

	// Part 1
	head, tail = PerformMoves(moves, head, tail)
	fmt.Println(len(tail.visited))

	// Part 2
	knots = InitSim(dim-1, 10)
	moves = ParseMoves(lines)
	knots = PerformMovesMany(knots, moves)
	tail = knots[len(knots)-1]
	fmt.Println(tail.visited)
	fmt.Println(len(tail.visited))
}
