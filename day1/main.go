package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

/*
--- Day 1: Trebuchet?! ---

[...]The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.
You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

[...]As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a
very young Elf.[...]

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that
the Elves now need to recover.

On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/
func main() {
	fmt.Println("hello World - Day 1")

	path := "input.txt"
	// Task 1 - Get the numbers
	file, err := os.Open(path)

	if err != nil {
		check(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {

		x := scanner.Text()
		word := strings.Split(x, "")

		// store numbers from each line
		myInt := []int{}

		for _, value := range word {
			if valueInt, err := strconv.Atoi(value); err == nil {
				myInt = append(myInt, valueInt)
			}
		} // end for - store numbers

		count = count + myInt[0]*10 + myInt[len(myInt)-1]

	} // end for - scanner

	// Task 2 - Sum
	fmt.Println(count)
}
