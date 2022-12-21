package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func ToIntArray(lines []string) []int {
	arr := make([]int, 0)
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		arr = append(arr, n)
	}
	return arr
}

func Mod(d, m int) int {
	var res int = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}

func IndexOf(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}

	return -1
}

func InsertInt(array []int, value int, index int) []int {
	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func RemoveInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func Mix(nums []int, times int) []int {
	l := len(nums)

	original := make([]int, l)
	copy(original, nums)

	indices := make([]int, l)
	for i := range nums {
		indices[i] = i
	}

	for t := 0; t < times; t++ {
		for i, v := range original {
			if v == 0 {
				continue
			}

			j := IndexOf(indices, i)
			x := indices[j]
			indices = RemoveInt(indices, j)
			k := Mod(j+v, l-1)
			indices = InsertInt(indices, x, k)
		}
	}

	result := make([]int, 0)
	for _, i := range indices {
		result = append(result, nums[i])
	}

	return result
}

func ApplyDecryptionKey(encryption []int, key int) []int {
	for i, n := range encryption {
		encryption[i] = n * key
	}
	return encryption
}

func main() {
	lines := ReadLines()

	// Part 1
	encryptedCoords := ToIntArray(lines)
	decrypted := Mix(encryptedCoords, 1)
	i := IndexOf(decrypted, 0)
	a := decrypted[(i+1000)%len(decrypted)]
	b := decrypted[(i+2000)%len(decrypted)]
	c := decrypted[(i+3000)%len(decrypted)]
	fmt.Println(a + b + c)

	// Part 2
	decryptionKey := 811589153
	encryptedCoords = ApplyDecryptionKey(encryptedCoords, decryptionKey)
	decrypted = Mix(encryptedCoords, 10)
	i = IndexOf(decrypted, 0)
	a = decrypted[(i+1000)%len(decrypted)]
	b = decrypted[(i+2000)%len(decrypted)]
	c = decrypted[(i+3000)%len(decrypted)]
	fmt.Println(a + b + c)
}
