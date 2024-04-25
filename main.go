package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <string>")
	}

	asciiChars, err := LoadAsciiChars("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error loading ASCII characters", err)
		return
	}
	// fmt.Printf("%s\n", asciiChars['h'])
	args := os.Args[1]
	arguments := strings.Split(args, "\\n")
	if args == "" {
		return
	} else if args == "\\n" {
		fmt.Println()
		return
	}

	for _, arg := range arguments {
		if arg == "" {
			fmt.Println()
		} else {
			printAsciiArt(arg, asciiChars) // Pass argument and map to printAsciiArt
		}
	}
}

func printAsciiArt(text string, asciiChars map[byte][]string) {
	for i := 0; i < 8; i++ {
		printLine(text, asciiChars, i) // call helper function for each line
		fmt.Println()
	}
}

func printLine(text string, asciiChars map[byte][]string, line int) {
	for _, char := range text {
		if char == '\n' {
			fmt.Println()
		} else {
			fmt.Print(asciiChars[byte(char)][line]) // Print the corresponding line from the map
		}
	}
}
