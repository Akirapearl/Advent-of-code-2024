package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

type Colors struct {
	green int
	red   int
	blue  int
}

func Sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func main() {
	fmt.Println("Day 2")

	path := "test.txt"
	// Task 1 - Get the numbers
	file, err := os.Open(path)

	if err != nil {
		check(err)
	}
	defer file.Close()

	// Set values for colors
	var TotalColors Colors

	TotalColors.red = 12
	TotalColors.green = 13
	TotalColors.blue = 14

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//count := 0

	for scanner.Scan() {
		x := scanner.Text()

		// Split by : to get each subset
		parts := strings.Split(x, ": ")
		gameIDPart := parts[0]
		subsetsPart := parts[1] // e.g., "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

		// Extract the game ID from "Game X"
		gameID := strings.TrimSpace(strings.TrimPrefix(gameIDPart, "Game "))
		fmt.Println("GAME ======= ", gameID)
		//fmt.Println(subsetsPart)

		// "Does at any point the elf show you more than 12 red, 13 green, or 14 blue cubes at one time? If not, the game is valid."

		// Get the amount - colour pair via regex
		re := regexp.MustCompile(`(\d+) (\w+)`)
		// Find all matches
		matches := re.FindAllStringSubmatch(subsetsPart, -1)

		myInt := []int{}
		var isWin bool
		count := 0
		for _, match := range matches {
			actualcount, err := strconv.Atoi(match[1])
			check(err)
			if !(actualcount >= TotalColors.red) && match[2] == "red" {
				isWin = true

			} else if !(actualcount >= TotalColors.green) && match[2] == "green" {
				isWin = true

			} else if !(actualcount >= TotalColors.blue) && match[2] == "blue" {
				isWin = true

			} else {
				isWin = false
				break
			}
		}
		fmt.Println(isWin)
		/*
			isGameID, err := strconv.Atoi(gameID)
			check(err)
			if isWin {
				myInt = append(myInt, isGameID)
				//fmt.Println(myInt)
			}
			fmt.Println(Sum(myInt))*/
		count = count + myInt[0]*10 + myInt[len(myInt)-1]
		fmt.Println(count)
	}

}
