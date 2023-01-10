package main

import (
	"bufio"
	"log"
	"os"
)

type directory struct {
	subdirectories []directory
	files          []file
}

type file struct {
	name string
	size int
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

	var resultP1, resultP2 int
	for fileScanner.Scan() {
		line := fileScanner.Text()
	}

	log.Printf("P1 : %d", resultP1)
	log.Printf("P2 : %d", resultP2)
}
