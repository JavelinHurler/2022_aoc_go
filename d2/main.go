package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type Token uint8

const (
	Rock     Token = 0
	Paper          = 1
	Scissors       = 2
)

type Result uint8

const (
	Loose Result = 0
	Draw         = 1
	Win          = 2
)

var opponentTokenMyTokenToScoreMap = map[Token]map[Token]uint32{
	Rock: map[Token]uint32{
		Rock:     1 + 3,
		Paper:    2 + 6,
		Scissors: 3 + 0,
	},
	Paper: map[Token]uint32{
		Rock:     1 + 0,
		Paper:    2 + 3,
		Scissors: 3 + 6,
	},
	Scissors: map[Token]uint32{
		Rock:     1 + 6,
		Paper:    2 + 0,
		Scissors: 3 + 3,
	},
}

var opponentTokenResultToScoreMap = map[Token]map[Result]uint32{
	Rock: map[Result]uint32{
		Loose: 3 + 0, //Scissors
		Draw:  1 + 3, //Rock
		Win:   2 + 6, //Paper
	},
	Paper: map[Result]uint32{
		Loose: 1 + 0, //Rock
		Draw:  2 + 3, //Paper
		Win:   3 + 6, //Scissors
	},
	Scissors: map[Result]uint32{
		Loose: 2 + 0, //Paper
		Draw:  3 + 3, //Scissors
		Win:   1 + 6, //Rock
	},
}

func getTokenFromLineABC(line string) (Token, error) {
	if strings.Contains(line, "A") {
		return Rock, nil
	} else if strings.Contains(line, "B") {
		return Paper, nil
	} else if strings.Contains(line, "C") {
		return Scissors, nil
	} else {
		return Rock, errors.New("Could not parse opponents token")
	}
}

func getTokenFromLineXYZ(line string) (Token, error) {
	if strings.Contains(line, "X") {
		return Rock, nil
	} else if strings.Contains(line, "Y") {
		return Paper, nil
	} else if strings.Contains(line, "Z") {
		return Scissors, nil
	} else {
		return Rock, errors.New("Could not parse my token")
	}
}

func getResultFromLineXYZ(line string) (Result, error) {
	if strings.Contains(line, "X") {
		return Loose, nil
	} else if strings.Contains(line, "Y") {
		return Draw, nil
	} else if strings.Contains(line, "Z") {
		return Win, nil
	} else {
		return Loose, errors.New("Could not parse my token")
	}
}

func getScoresForRound(line string) (scoreP1 uint32, scoreP2 uint32) {
	opponentToken, err := getTokenFromLineABC(line)
	if err != nil {
		log.Fatal("getTokenFromLineABC : #v", err)
		os.Exit(-1)
	}

	myToken, err := getTokenFromLineXYZ(line)
	if err != nil {
		log.Fatal("getTokenFromLineXYZ : #v", err)
		os.Exit(-1)
	}

	expectedResult, err := getResultFromLineXYZ(line)
	if err != nil {
		log.Fatal("getResultFromLineXYZ : #v", err)
		os.Exit(-1)
	}

	scoreP1 = opponentTokenMyTokenToScoreMap[opponentToken][myToken]
	scoreP2 = opponentTokenResultToScoreMap[opponentToken][expectedResult]

	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("os.Open: %s", err.Error())
		os.Exit(-1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var gameScoreP1 uint32 = 0
	var gameScoreP2 uint32 = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		roundScoreP1, roundScoreP2 := getScoresForRound(line)
		gameScoreP1 += roundScoreP1
		gameScoreP2 += roundScoreP2
	}

	log.Printf("P1 : %d", gameScoreP1)
	log.Printf("P2 : %d", gameScoreP2)
}
