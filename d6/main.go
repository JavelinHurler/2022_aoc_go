package main

import (
	"bufio"
	"errors"
	"log"
	"os"
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

	loop := 0
	var resultP1, resultP2 int
	for fileScanner.Scan() {
		loop += 1
		line := fileScanner.Text()

		resultP1, err = getIndexOfFirstMarker(line, 4)
		if err != nil {
			log.Fatal("Call1 | ", err.Error())
		}

		resultP2, err = getIndexOfFirstMarker(line, 14)
		if err != nil {
			log.Fatal("Call2 | ", err.Error())
		}
	}

	if loop > 1 {
		log.Fatal("More then one line in file")
	}

	log.Printf("P1 : %d", resultP1)
	log.Printf("P2 : %d", resultP2)
}

func getIndexOfFirstMarker(line string, count int) (int, error) {
	for i := (count - 1); i < len(line); i += 1 {
		if isMarker(line[i-count+1 : i+1]) {
			return i + 1, nil
		}
	}
	return 0, errors.New("Error occured in \"getIndexOfFirstMarker()\" : no marker in string")
}

func isMarker(subline string) bool {
	for i := 0; i < len(subline); i += 1 {
		for j := (i + 1); j < len(subline); j += 1 {
			if subline[i] == subline[j] {
				return false
			}
		}
	}
	return true
}
