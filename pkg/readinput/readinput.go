package readinput

import (
	"io"
	"os"
	"strings"
)

func ReadInput(fileName string) []string {
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
