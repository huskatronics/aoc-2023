package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCardString(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := Card{winningNum: []string{"41", "48", "83", "86", "17"}, ownNum: []string{"83", "86", "6", "31", "17", "9", "48", "53"}}

	result := ParseCardString(input)
	assert.ElementsMatch(t, expected.winningNum, result.winningNum)
	assert.ElementsMatch(t, expected.ownNum, result.ownNum)
}

func TestGetWinningNumCount(t *testing.T) {
	input := Card{winningNum: []string{"41", "48", "83", "86", "17"}, ownNum: []string{"83", "86", "6", "31", "17", "9", "48", "53"}}
	expect := 4

	result := GetWinningNumCount(input)
	assert.Equal(t, expect, result)

}

func TestCalculateFinalScore(t *testing.T) {
	input := 4
	expect := 8

	result := CalculateFinalScore(input)
	assert.Equal(t, expect, result)
}
