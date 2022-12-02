package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PlayerMove struct {
	id          string
	counters    string
	counteredBy string
	equivalent  string
	value       int
}

var Moves = map[string]PlayerMove{
	"A": {
		id:          "Rock",
		counters:    "Z",
		counteredBy: "Y",
		equivalent:  "X",
		value:       1,
	},
	"B": {
		id:          "Paper",
		counters:    "X",
		counteredBy: "Z",
		equivalent:  "Y",
		value:       2,
	},
	"C": {
		id:          "Scissors",
		counters:    "Y",
		counteredBy: "X",
		equivalent:  "Z",
		value:       3,
	},
}

const (
	Win  = "Z"
	Draw = "Y"
	Lose = "X"
)

const (
	WinScore  = 6
	DrawScore = 3
)

func HasWon(player1 string, player2 string) bool {
	return player2 == Moves[player1].counteredBy
}

func HasDrawn(player1 string, player2 string) bool {
	return player2 == Moves[player1].equivalent
}

func FindMove(player string) (PlayerMove, bool) {
	for _, v := range Moves {
		if v.equivalent == player {
			return v, true
		}
	}

	return PlayerMove{}, false
}

func GetResponse(player string, strategy string) string {
	if strategy == Win {
		return Moves[player].counteredBy
	} else if strategy == Draw {
		return Moves[player].equivalent
	}
	return Moves[player].counters
}

func CountTotalScores() (int, int) {
	totalScore := 0
	strategicScore := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		round := scanner.Text()
		players := strings.Split(round, " ")

		if move, found := FindMove(players[1]); found {
			totalScore += move.value
		}

		if HasWon(players[0], players[1]) {
			totalScore += WinScore
		} else if HasDrawn(players[0], players[1]) {
			totalScore += DrawScore
		}

		response := GetResponse(players[0], players[1])
		if move, found := FindMove(response); found {
			strategicScore += move.value
		}

		if HasWon(players[0], response) {
			strategicScore += WinScore
		} else if HasDrawn(players[0], response) {
			strategicScore += DrawScore
		}
	}

	return totalScore, strategicScore
}

func main() {
	totalScore, strategicScore := CountTotalScores()
	fmt.Println(totalScore)     // Part 1
	fmt.Println(strategicScore) // Part 2
}
