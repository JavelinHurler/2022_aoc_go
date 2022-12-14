package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type VecUint32 []uint32

func (arr VecUint32) Len() int {
	return len(arr)
}

func (arr VecUint32) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr VecUint32) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("os.Open %s", err.Error())
		os.Exit(-1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var currentCalories uint32
	caloriesPerReindeer := VecUint32{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			caloriesPerReindeer = append(caloriesPerReindeer, currentCalories)
			currentCalories = 0
		} else {
			lineInt, err := strconv.ParseUint(line, 10, 32)
			if err != nil {
				fmt.Println("strconv.ParseUint %s", err.Error())
				os.Exit(-1)
			}
			currentCalories += uint32(lineInt)
		}
	}

	sort.Sort(caloriesPerReindeer)
	lastIndex := caloriesPerReindeer.Len() - 1

	sumTopThreeCalories := caloriesPerReindeer[lastIndex] + caloriesPerReindeer[lastIndex-1] + caloriesPerReindeer[lastIndex-2]

	fmt.Println("P1", caloriesPerReindeer[lastIndex])
	fmt.Println("P2", sumTopThreeCalories)
}
