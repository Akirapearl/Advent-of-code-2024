package main

/*
--- Part Two ---

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?

-- 53894?
*/
import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Numbers struct {
	Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine string
}

// Initializes Numbers struct with spelled-out numbers
func NewNumbers() Numbers {
	return Numbers{
		Zero:  "zero",
		One:   "one",
		Two:   "two",
		Three: "three",
		Four:  "four",
		Five:  "five",
		Six:   "six",
		Seven: "seven",
		Eight: "eight",
		Nine:  "nine",
	}
}

// Inspired by https://github.com/prathamesh-88/advent-of-code-2023/blob/master/Day01/trebuchet2.go
func extractDigitsAndSpeltStrings(str string, numbers Numbers) []int {
	digits := []int{}

	for i, char := range str {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		} else {
			// Use reflection to iterate over Numbers struct fields
			numbersValue := reflect.ValueOf(numbers)
			for j := 0; j < numbersValue.NumField(); j++ {
				spelling := numbersValue.Field(j).String()
				if strings.HasPrefix(str[i:], spelling) {
					digits = append(digits, j) // `j` represents the number
					break
				}
			}
		}
	}

	return digits
}

func main() {
	fmt.Println("hello World - Day 1")

	path := "input.txt"
	file, err := os.Open(path)
	if err != nil {
		check(err)
	}
	defer file.Close()

	// Initialize Numbers struct
	numbers := NewNumbers()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		x := scanner.Text()

		// Extract digits and spelled-out numbers using extractDigitsAndSpeltStrings
		myInt := extractDigitsAndSpeltStrings(x, numbers)

		// Sum up the values
		if len(myInt) > 0 {
			count += myInt[0]*10 + myInt[len(myInt)-1]
		}
	}

	fmt.Println("Sum:", count)
}
