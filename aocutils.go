// Various bits of reused code for Advent of Code for Tim
package aocutils

import (
	"bufio"
	"log"
	"os"
)

// Returns the word "Hello.".  This was, and is, a test.
func Hello() string {
	return "Hello."
}

// Return the min and max values of an array of ints
func MinMaxInt(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

// Reads a file from the current directory and packs it into a array of strings.
func Readfile(filename string) []string {
	dir, _ := os.Getwd()

	file, err := os.Open(dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var commands []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands
}
