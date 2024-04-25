package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadAsciiChars(filename string) (map[byte][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	asciiChars := make(map[byte][]string)

	scanner := bufio.NewScanner(file)

	currentChar := byte(' ')
	count := 0
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if count != 8 {
			asciiChars[currentChar] = append(asciiChars[currentChar], line)
			count++
		} else {
			currentChar++
			count = 0
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return asciiChars, nil
}
