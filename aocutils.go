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
