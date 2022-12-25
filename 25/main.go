package main

import (
	"bufio"
	"fmt"
	"os"
)

var DigitsLookup = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
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

func SnafuToDecimal(snafu string) (n int) {
	multiple := 1
	for i := len(snafu) - 1; i >= 0; i-- {
		if i < len(snafu)-1 {
			multiple *= 5
		}

		ch := rune(snafu[i])
		n += DigitsLookup[ch] * multiple
	}
	return n
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

func DecimalToSnafu(n int) (snafu string) {
	if n == 0 {
		return ""
	}

	q, r := divmod(n+2, 5)
	chars := "=-012"
	return DecimalToSnafu(q) + string(chars[r])
}

func CalcSnafuSum(lines []string) (sum int) {
	for _, snafu := range lines {
		sum += SnafuToDecimal(snafu)
	}
	return sum
}

func main() {
	lines := ReadLines()
	snafuSum := CalcSnafuSum(lines)
	snafu := DecimalToSnafu(snafuSum)
	fmt.Println(snafu)
}
