package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
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

	// Replace with your path to the actual input file
	path := "input.txt"
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

	var validGameIDs []int
	for scanner.Scan() {
		x := scanner.Text()

		// Split by ": " to get each subset
		parts := strings.Split(x, ": ")
		gameIDPart := parts[0]
		subsetsPart := parts[1] // e.g., "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

		// Extract the game ID from "Game X"
		gameID := strings.TrimSpace(strings.TrimPrefix(gameIDPart, "Game "))
		//fmt.Println("GAME ======= ", gameID)
		//fmt.Println(subsetsPart)

		// "Does at any point the elf show you more than 12 red, 13 green, or 14 blue cubes at one time? If not, the game is valid."

		// Get the amount-color pair via regex
		re := regexp.MustCompile(`(\d+) (\w+)`)
		// Find all matches
		matches := re.FindAllStringSubmatch(subsetsPart, -1)

		isValid := true
		for _, match := range matches {
			actualCount, err := strconv.Atoi(match[1])
			check(err)

			// If the actual count exceeds the threshold for any color, mark the game as invalid
			if match[2] == "red" && actualCount > TotalColors.red {
				isValid = false
				break
			} else if match[2] == "green" && actualCount > TotalColors.green {
				isValid = false
				break
			} else if match[2] == "blue" && actualCount > TotalColors.blue {
				isValid = false
				break
			}
		}

		// If the game is valid, add its ID to the list
		if isValid {
			gameIDInt, err := strconv.Atoi(gameID)
			check(err)
			validGameIDs = append(validGameIDs, gameIDInt)
		}
	}

	// Output the sum of valid game IDs after processing all lines
	fmt.Println(Sum(validGameIDs))
}
