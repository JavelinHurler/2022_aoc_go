package main

import (
	"bufio"
	"log"
	"os"

	"example.com/2022_aoc_d3/utils"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("os.Open: %s", err.Error())
		os.Exit(-1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var (
		last3lines                   [3]string
		lineNumber, scoreP1, scoreP2 uint32
	)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		last3lines[lineNumber%3] = line

		commonLowerCaseChar, commonUpperCaseChar, err := utils.GetCommonUpperAndLowerChar(line)
		if err != nil {
			log.Fatalf("seperateLine: %s", err.Error())
		}

		lowerScore, err := utils.MapCharToScore(commonLowerCaseChar)
		if err != nil {
			log.Fatalf("mapCharToScore: %s", err.Error())
		}

		upperScore, err := utils.MapCharToScore(commonUpperCaseChar)
		if err != nil {
			log.Fatalf("mapCharToScore: %s", err.Error())
		}

		scoreP1 += (uint32(lowerScore) + uint32(upperScore))

		lineNumber += 1

		if lineNumber%3 == 0 && lineNumber != 0 {
			groupChar, err := utils.GetCommonCharGroup(last3lines[0], last3lines[1], last3lines[2])
			if err != nil {
				log.Fatalf("getCommonChar: %s", err.Error())
			}

			groupScore, err := utils.MapCharToScore(groupChar)
			if err != nil {
				log.Fatalf("mapCharToScore: %s", err.Error())
			}

			scoreP2 += uint32(groupScore)
		}
	}

	log.Printf("P1 : %d", scoreP1)
	log.Printf("P2 : %d", scoreP2)
}
