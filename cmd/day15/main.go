package main

import (
	"aoc2023-golang/pkg/readinput"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := readinput.ReadInput("input.txt", ",")
	result1 := solve1(data)
	fmt.Printf("Answer for part 1: %d\n", result1)
	result2 := solve2(data)
	fmt.Printf("Answer for part 2: %d\n", result2)
}

func solve1(data []string) int {
	totalSum := 0
	for _, item := range data {
		value := getHashValue(item)
		totalSum += value
	}
	return totalSum
}

type lensType struct {
	label string
	focal int
}

func solve2(data []string) int {
	boxData := [256][]lensType{}
	for _, item := range data {
		label, focal, isAddition := processOrder(item)
		boxNum := getHashValue(label)
		currentBox := boxData[boxNum]
		if isAddition {
			if len(currentBox) == 0 {
				currentBox = append(currentBox, lensType{label: label, focal: focal})
				boxData[boxNum] = currentBox
				continue
			}
			for i, lens := range currentBox {
				if lens.label == label {
					currentBox[i] = lensType{label: label, focal: focal}
					break
				}
				if i == len(currentBox)-1 {
					currentBox = append(currentBox, lensType{label: label, focal: focal})
					boxData[boxNum] = currentBox
				}
			}
		} else {
			if len(currentBox) == 0 {
				continue
			}
			for i, lens := range currentBox {
				if lens.label == label {
					currentBox = slices.Delete(currentBox, i, i+1)
					boxData[boxNum] = currentBox
					break
				}
			}
		}
	}
	return getFinalSum(boxData)
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

func processOrder(order string) (string, int, bool) {
	result := [2]string{}
	isAddition := false
	focal := 0

	if strings.IndexRune(order, '=') != -1 {
		isAddition = true
		parsed := strings.Split(order, "=")
		result[0] = parsed[0]
		result[1] = parsed[1]
	} else if strings.IndexRune(order, '-') != -1 {
		parsed := strings.Split(order, "-")
		result[0] = parsed[0]
		result[1] = "0"
	}

	label := result[0]
	focalValue, err := strconv.ParseInt(result[1], 10, 64)
	if err != nil {
		panic(err)
	}
	focal = int(focalValue)

	return label, focal, isAddition

}

func getFinalSum(boxData [256][]lensType) int {
	sum := 0
	for i, box := range boxData {
		if len(box) == 0 {
			continue
		}
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focal
		}
	}
	return sum
}
