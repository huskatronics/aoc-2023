package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type itemType struct {
	line, startIdx, endIdx, value int
}

type locator struct {
	line, idx int
}

var lastLine int

func main() {
	data := readInput("input.txt")
	answerOne := solve1(data)
	fmt.Printf("Answer for part 1: %d \n", answerOne)
	answerTwo := solve2(data)
	fmt.Printf("Answer for part 2: %d \n", answerTwo)
}

func readInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	input, err := io.ReadAll(file)
	trimmedInput := strings.TrimSpace(string(input))
	if err != nil {
		panic(err)
	}
	return strings.Split(trimmedInput, "\r\n")
}

var symbolRegexp *regexp.Regexp = regexp.MustCompile("[^a-zA-Z0-9.]")

func solve1(data []string) int {
	total := 0
	lastLine = len(data) - 1

	var digitList []itemType
	var symbolList []locator
	for i, line := range data {
		normalisedLine := symbolRegexp.ReplaceAllString(line, "*")
		digits, symbols := parseLine(i, normalisedLine)
		for _, item := range digits {
			digitList = append(digitList, item)
		}
		for _, item := range symbols {
			symbolList = append(symbolList, item)
		}

	}

	fmt.Printf("locations: %+v \n", symbolList)
	for _, item := range digitList {
		if isPart(item, symbolList) {
			fmt.Printf("%+v is a part\n", item)
			total += item.value
		}
	}

	return total
}
func solve2(data []string) int {
	total := 0
	return total
}

func isSymbolChar(char byte) bool {
	alphabetAndFriends := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz."
	return !slices.Contains([]byte(alphabetAndFriends), char)
}

var digitsRegexp *regexp.Regexp = regexp.MustCompile("\\d+")

func isDigits(item string) bool {
	return digitsRegexp.MatchString(item)
}

var dotRegexp *regexp.Regexp = regexp.MustCompile("\\.+")

func parseLine(lineNum int, line string) ([]itemType, []locator) {
	var digits []itemType
	var symbols []locator
	paddedString := symbolRegexp.ReplaceAllString(line, ".*.")
	cleanLine := dotRegexp.ReplaceAllString(paddedString, ".")
	items := strings.Split(cleanLine, ".")
	searchIndex := 0
	for _, item := range items {
		if isDigits(item) {
			value, _ := strconv.ParseInt(item, 10, 64)
			startIdx := strings.Index(line[searchIndex:], item) + searchIndex
			endIdx := startIdx + len(item) - 1
			searchIndex = endIdx
			digits = append(digits, itemType{line: lineNum, value: int(value), startIdx: startIdx, endIdx: endIdx})
			continue
		}
		for _, char := range item {
			if isSymbolChar(byte(char)) {
				index := strings.Index(line[searchIndex:], string(char))
				symbols = append(symbols, locator{line: lineNum, idx: index + searchIndex})
				searchIndex = index + 1
			}
		}
	}
	return digits, symbols
}

func isPart(item itemType, locators []locator) bool {
	return hasNeighbouringLocator(item.line, item.startIdx, item.endIdx, locators)
}

func hasNeighbouringLocator(lineNum, startIndex, endIndex int, locators []locator) bool {
	for _, locator := range locators {
		if locator.line == lineNum {
			if locator.idx == startIndex-1 || locator.idx == endIndex+1 {
				return true
			}
		} else if lineNum == 0 {
			if locator.line == lineNum+1 {
				if locator.idx >= startIndex-1 && locator.idx <= endIndex+1 {
					return true
				}
			}
		} else if lineNum == lastLine {
			if locator.line == lineNum-1 {
				if locator.idx >= startIndex-1 && locator.idx <= endIndex+1 {
					return true
				}
			}
		} else if locator.line == lineNum-1 || locator.line == lineNum+1 {
			if locator.idx >= startIndex-1 && locator.idx <= endIndex+1 {
				return true
			}
		}
	}
	return false
}
