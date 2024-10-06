package main

import (
	"fmt"
	"log"
	"os"
)

func readFile() ([]byte, error) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		return nil, fmt.Errorf("reading error: %w", err)
	}
	return input, nil
}

func main() {
	file, err := readFile()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(problemOne(file))
	fmt.Println(problemTwo(file))
}

func problemOne(f []byte) int {
	var floor int
	for _, v := range f {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func problemTwo(f []byte) uint {
	var floor int
	var position uint
	for _, v := range f {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
		position++
		if floor < 0 {
			break
		}
	}
	return position
}
