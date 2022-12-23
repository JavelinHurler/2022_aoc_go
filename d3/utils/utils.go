package utils

import (
	"errors"
	"unicode"

	"example.com/2022_aoc_d3/set"
)

func getCommonChar(line1, line2 string) byte {
	lineSet1 := make(set.Set)
	lineSet1.FillFromString(line1)

	lineSet2 := make(set.Set)
	lineSet2.FillFromString(line2)

	for char := range lineSet1 {
		if lineSet2.ValueExists(char) {
			return char
		}
	}

	return 0
}

func GetCommonCharGroup(line1, line2, line3 string) (byte, error) {
	lineSet1 := make(set.Set)
	lineSet1.FillFromString(line1)

	lineSet2 := make(set.Set)
	lineSet2.FillFromString(line2)

	lineSet3 := make(set.Set)
	lineSet3.FillFromString(line3)

	for char := range lineSet1 {
		if lineSet2.ValueExists(char) && lineSet3.ValueExists(char) {
			return char, nil
		}
	}
	return 0, errors.New("no common char")
}

func seperateStringByCase(line string) (lowerPart string, upperPart string, err error) {
	for _, char := range line {
		if !unicode.IsLetter(char) {
			err = errors.New("String contains non letter unicode symbol")
			return
		}
		if unicode.IsUpper(char) {
			upperPart += string(char)
		} else {
			lowerPart += string(char)
		}
	}
	return
}

func seperateLine(line string) (string, string, error) {
	length := len(line)
	if length%2 != 0 {
		return "", "", errors.New("Length of line is odd")
	}

	middle := length / 2

	firstHalf := line[0:middle]
	secondHalf := line[middle:length]

	return firstHalf, secondHalf, nil
}

func GetCommonUpperAndLowerChar(line string) (byte, byte, error) {
	firstHalf, secondHalf, err := seperateLine(line)
	if err != nil {
		return 0, 0, err
	}

	firstHalfLower, firstHalfUpper, err := seperateStringByCase(firstHalf)
	if err != nil {
		return 0, 0, err
	}

	secondHalfLower, secondHalfUpper, err := seperateStringByCase(secondHalf)
	if err != nil {
		return 0, 0, err
	}

	commonLowerCaseChar := getCommonChar(firstHalfLower, secondHalfLower)
	commonUpperCaseChar := getCommonChar(firstHalfUpper, secondHalfUpper)

	return commonLowerCaseChar, commonUpperCaseChar, nil
}

func MapCharToScore(char byte) (byte, error) {
	if char == 0 {
		return 0, nil
	}
	if char >= 97 && char <= 122 {
		return char - 96, nil
	}
	if char >= 65 && char <= 90 {
		return char - 38, nil
	}
	return 0, errors.New("Char to be mapped is not an letter")
}
