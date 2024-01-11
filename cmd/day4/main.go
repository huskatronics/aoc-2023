package main

import (
	"aoc2023-golang/pkg/common"
	"aoc2023-golang/pkg/readinput"
	"fmt"
	"slices"
	"strings"
)

func main() {
	data := readinput.ReadInput("input.txt", "\r\n")
	res1 := Solve1(data)
	fmt.Printf("Answer for the first part: %d\n", res1)
	res2 := Solve2(data)
	fmt.Printf("Answer for the second part: %d\n", res2)
}

func Solve1(data []string) int {
	result := 0
	for _, card := range data {
		parsedCard := ParseCardString(card)
		winningCount := GetWinningNumCount(parsedCard)
		if winningCount > 0 {
			result += CalculateFinalScore(winningCount)
		}
	}
	return result
}

func Solve2(data []string) int {
	result := 0
	return result
}

type Card struct {
	winningNum, ownNum []string
}

func ParseCardString(cardString string) Card {
	card := strings.Split(cardString, ":")
	cardDetail := strings.TrimSpace(card[1])
	numberLists := strings.Split(cardDetail, "|")
	winningNums := strings.Split(strings.TrimSpace(numberLists[0]), " ")
	cleanWinningNums := slices.DeleteFunc(winningNums, func(item string) bool { return item == "" })
	ownNums := strings.Split(strings.TrimSpace(numberLists[1]), " ")
	cleanOwnNums := slices.DeleteFunc(ownNums, func(item string) bool { return item == "" })
	return Card{ownNum: cleanOwnNums, winningNum: cleanWinningNums}
}

func GetWinningNumCount(card Card) int {
	count := 0
	for _, num := range card.winningNum {
		if common.IsStringInArray(card.ownNum, num) {
			count += 1
		}
	}
	return count
}

func CalculateFinalScore(winningCount int) int {
	return 1 << (winningCount - 1)
}
