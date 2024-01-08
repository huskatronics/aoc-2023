package readinput

import (
	"io"
	"os"
	"strings"
)

func ReadInput(fileName string, separator string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	input, err := io.ReadAll(file)
	trimmedInput := strings.TrimSpace(string(input))
	if err != nil {
		panic(err)
	}
	parsedInput := strings.Split(trimmedInput, separator)

	var cleanInput []string
	for _, item := range parsedInput {
		cleanInput = append(cleanInput, strings.TrimSpace(item))
	}
	return cleanInput
}
