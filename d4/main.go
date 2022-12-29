package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type slice struct {
	begin uint32
	end   uint32
}

func (this slice) contains(that slice) bool {
	return (this.begin <= that.begin && this.end >= that.end)
}

func (this slice) isIn(that slice) bool {
	return ((this.begin >= that.begin && this.begin <= that.end) || (this.end >= that.begin && this.end <= that.end))
}

func (this *slice) parseFromString(input string) error {
	var begin, end uint32
	_, err := fmt.Sscanf(input, "%d-%d", &begin, &end)
	if err != nil {
		return err
	}
	this.begin = begin
	this.end = end
	return nil
}

func parseLine(line string) (slice, slice, error) {
	var begin, end slice

	splits := strings.Split(line, ",")
	if len(splits) != 2 {
		return begin, end, errors.New("Line contained " + strconv.Itoa(len(splits)) + " commas instead of one")
	}

	err := begin.parseFromString(splits[0])
	if err != nil {
		return begin, end, err
	}

	err = end.parseFromString(splits[1])
	if err != nil {
		return begin, end, err
	}

	return begin, end, nil
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

	var scoreP1, scoreP2 uint32

	for fileScanner.Scan() {
		line := fileScanner.Text()
		begin, end, err := parseLine(line)
		if err != nil {
			log.Fatal("error while parsing line")
		}

		if begin.contains(end) || end.contains(begin) {
			scoreP1 += 1
		}

		if begin.isIn(end) || end.isIn(begin) {
			scoreP2 += 1
		}
	}

	log.Printf("P1 : %d", scoreP1)
	log.Printf("P2 : %d", scoreP2)
}
