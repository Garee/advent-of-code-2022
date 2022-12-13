package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

func ParsePackets(lines []string) []any {
	packets := make([]any, 0)

	for i := 0; i < len(lines); i += 3 {
		var p1, p2 any
		json.Unmarshal([]byte(lines[i]), &p1)
		json.Unmarshal([]byte(lines[i+1]), &p2)
		packets = append(packets, p1)
		packets = append(packets, p2)
	}

	return packets
}

func Compare(p1, p2 any) int {
	a, okA := p1.(float64)
	b, okB := p2.(float64)
	if okA && okB {
		return int(a) - int(b)
	}

	var fList []any
	var sList []any

	switch p1.(type) {
	case []any, []float64:
		fList = p1.([]any)
	case float64:
		fList = []any{p1}
	}

	switch p2.(type) {
	case []any, []float64:
		sList = p2.([]any)
	case float64:
		sList = []any{p2}
	}

	for i := range fList {
		if len(sList) <= i {
			return 1
		}

		if c := Compare(fList[i], sList[i]); c != 0 {
			return c
		}
	}

	if len(sList) == len(fList) {
		return 0
	}

	return -1
}

func FindInOrder(packets []any) []int {
	indices := make([]int, 0)
	for i := 0; i < len(packets); i += 2 {
		var p1, p2 = packets[i], packets[i+1]
		if c := Compare(p1, p2); c <= 0 {
			indices = append(indices, (i/2)+1)
		}
	}
	return indices
}

func Sum(arr []int) (sum int) {
	for _, n := range arr {
		sum += n
	}
	return sum
}

func CorrectOrdering(packets []any) []any {
	var a, b any
	json.Unmarshal([]byte("[[2]]"), &a)
	json.Unmarshal([]byte("[[6]]"), &b)
	packets = append(packets, a, b)

	sort.Slice(packets, func(i, j int) bool {
		return Compare(packets[i], packets[j]) < 0
	})

	return packets
}

func FindDividers(packets []any) (int, int) {
	a, b := -1, -1
	for i, packet := range packets {
		s, _ := json.Marshal(packet)
		if string(s) == "[[2]]" {
			a = i + 1
		} else if string(s) == "[[6]]" {
			b = i + 1
		}
	}
	return a, b
}

func main() {
	lines := ReadLines()
	packets := ParsePackets(lines)

	// Part 1
	indices := FindInOrder(packets)
	fmt.Println(Sum(indices))

	// Part 2
	correct := CorrectOrdering(packets)
	a, b := FindDividers(correct)
	fmt.Println(a * b)
}
