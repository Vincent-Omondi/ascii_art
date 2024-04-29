package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage: go run . <string>, use string as one argument in quotes")
		return
	}

	asciiChars, err := LoadAsciiChars("standard.txt")
	if err != nil {
		fmt.Println("Error loading ASCII characters", err)
		return
	}

	args := os.Args[1]

	args = strings.ReplaceAll(args, "\n", "\\n")
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
	text = strings.ReplaceAll(text, "\\r", "\r")
	text = strings.ReplaceAll(text, "\\b", "\b")
	text = strings.ReplaceAll(text, "\\t", "    ")
	text = strings.ReplaceAll(text, "\\f", "\f")
	text = strings.ReplaceAll(text, "\\a", "\a")
	text = strings.ReplaceAll(text, "\\v", "\v")
	for _, char := range text {
		if char > 127 || char < 32 {
			fmt.Printf("Error: Character %q is not accepted\n", char)
			return
		}
	}
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
