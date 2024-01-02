package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re1 *regexp.Regexp = regexp.MustCompile(`\d`)
var numbers []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, _ := os.Open("input.txt")
	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	parsedInput := parse(string(input))
	//result1 := solve1(parsedInput)
	//fmt.Printf("Answer for part one: %d", result1)
	result2 := solve2(parsedInput)
	fmt.Printf("Answer for part two: %d \n", result2)
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func solve1(input []string) int {
	sum := 0
	for _, line := range input {
		digits := re1.FindAllString(line, -1)
		fmt.Println(digits)
		if len(digits) == 0 {
			continue
		}
		stringNumber := fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
		number, _ := strconv.ParseInt(stringNumber, 10, 64)
		sum += int(number)
	}
	return sum
}
func solve2(input []string) int {
	sum := 0
	for _, line := range input {
		var digits []string
		for i := 0; i < len(line); i++ {
			substring := line[i:]
			for _, number := range numbers {
				if strings.HasPrefix(substring, number) {
					digits = append(digits, number)
					break
				}
			}
		}
		if len(digits) == 0 {
			continue
		}
		parsedDigits := parseLineDigits(digits)
		stringNumber := fmt.Sprintf("%s%s", parsedDigits[0], parsedDigits[len(parsedDigits)-1])
		number, _ := strconv.ParseInt(stringNumber, 10, 64)
		sum += int(number)
	}
	return sum
}

func parseWordToDigit(input string) string {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return input
	}
}

func parseLineDigits(input []string) []string {
	var result []string
	for _, obj := range input {
		result = append(result, parseWordToDigit(obj))
	}
	return result
}
