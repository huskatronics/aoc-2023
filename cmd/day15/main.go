package main

import (
	"aoc2023-golang/pkg/readinput"
	"fmt"
)

func main() {
	data := readinput.ReadInput("input.txt", ",")
	result1 := solve1(data)
	fmt.Printf("Answer for part 1: %d\n", result1)
}

func solve1(data []string) int {
	totalSum := 0
	for _, item := range data {
		value := getHashValue(item)
		totalSum += value
	}
	return totalSum
}

func getHashValue(word string) int {
	total := 0
	for _, character := range word {
		//making sure it's utf-8
		if int(character) > 255 {
			continue
		}
		total += int(character)
		total *= 17
		total = total % 256
	}
	return total
}
