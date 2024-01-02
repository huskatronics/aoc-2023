package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type play struct {
	red, green, blue int
}

type game struct {
	id               int
	red, green, blue []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	input, _ := io.ReadAll(file)
	parsedInput := strings.Split(string(input), "\r\n")
	partOne := solve1(parsedInput)
	fmt.Printf("Answer for part 1: %d \n", partOne)
	partTwo := solve2(parsedInput)
	fmt.Printf("Answer for part 2: %d \n", partTwo)
}

func parsePlay(playString string) play {
	play := play{red: 0, green: 0, blue: 0}
	colours := strings.Split(playString, ", ")
	for _, item := range colours {
		result := strings.Split(item, " ")
		amount, _ := strconv.ParseInt(result[0], 10, 8)
		colour := result[1]
		switch colour {
		case "red":
			play.red = int(amount)
		case "blue":
			play.blue = int(amount)
		case "green":
			play.green = int(amount)
		}
	}
	return play
}

func parseGame(gameString string) game {
	var red []int
	var blue []int
	var green []int
	gameDetail := strings.Split(gameString, ": ")
	plays := strings.Split(gameDetail[1], "; ")
	for _, play := range plays {
		parsedPlay := parsePlay(play)
		red = append(red, parsedPlay.red)
		blue = append(blue, parsedPlay.blue)
		green = append(green, parsedPlay.green)
	}
	gameIdentifier := strings.Split(gameDetail[0], " ")
	gameId, _ := strconv.ParseInt(gameIdentifier[1], 10, 8)

	return game{id: int(gameId), red: red, green: green, blue: blue}
}

func solve1(input []string) int {
	sum := 0
	supply := play{red: 12, green: 13, blue: 14}
	for _, game := range input {
		if len(game) == 0 {
			continue
		}
		parsedGame := parseGame(game)
		maxRed := slices.Max(parsedGame.red)
		maxBlue := slices.Max(parsedGame.blue)
		maxGreen := slices.Max(parsedGame.green)
		if maxRed <= supply.red && maxGreen <= supply.green && maxBlue <= supply.blue {
			fmt.Println(parsedGame.id)
			fmt.Printf("Raw game data: %s\n", game)
			fmt.Printf("Required balls for red: %d, blue: %d, green: %d \n", maxRed, maxBlue, maxGreen)
			sum += parsedGame.id
		}
	}
	return sum
}

func solve2(input []string) int {
	sum := 0
	for _, game := range input {
		if len(game) == 0 {
			continue
		}
		parsedGame := parseGame(game)
		minRed := slices.Max(parsedGame.red)
		minBlue := slices.Max(parsedGame.blue)
		minGreen := slices.Max(parsedGame.green)
		sum += (minRed * minBlue * minGreen)
	}
	return sum
}
