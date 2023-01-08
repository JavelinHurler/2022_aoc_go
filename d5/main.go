package main

import (
	"bufio"
	"log"
	"os"

	"example.com/2022_aoc_d5/parser"
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

	var fofiParser, lofiParser parser.Parser
	fofiParser.New(parser.FoFi)
	lofiParser.New(parser.LoFi)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		err := fofiParser.ParseLine(line)
		if err != nil {
			log.Fatal("Error Occurred %s", err.Error())
		}

		err = lofiParser.ParseLine(line)
		if err != nil {
			log.Fatal("Error Occurred %s", err.Error())
		}
	}

	resultP1, err := fofiParser.Done()
	resultP2, err := lofiParser.Done()

	log.Printf("P1 : %s", resultP1)
	log.Printf("P2 : %s", resultP2)
}
