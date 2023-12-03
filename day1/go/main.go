package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"strconv"
	"time"
	)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var strToNum = make(map[string]int)

func init() {
	// Fill mapping of string to numbers
	for i, str := range numbers {
		strToNum[str] = i + 1
	}
}

func timer() func() {
	start := time.Now()
	return func() {
		log.Printf("Timer: %v", time.Since(start))
	}
}

func main() {
	defer timer()()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var accumPartOne int = 0
	var accumPartTwo int = 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Print("Reading line: ", line)
		numberPartOne := extractNumberPartOne(line)
		numberPartTwo := extractNumberPartTwo(line)
		accumPartOne = accumPartOne + numberPartOne
		accumPartTwo = accumPartTwo + numberPartTwo
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file", err)
	}

	log.Print("----------------")
	log.Print("Result part one: ", accumPartOne)
	log.Print("Result part two: ", accumPartTwo)
}

func extractNumberPartOne(line string) int {
	r := regexp.MustCompile(`\d`)
	matches := r.FindAllString(line, -1)

	first := matches[0]
	last := matches[len(matches) - 1]
	numberString := first + last
	number, err := strconv.Atoi(numberString)

	if err != nil {
		log.Fatal("Error extracting number from ", numberString, err)
	}
	log.Print("Part 1: ", number)
	return number
}


func extractNumberPartTwo(line string) int {
	numbersPattern := strings.Join(numbers, "|")
	stringForwards := fmt.Sprintf(`(\d|%s)`, numbersPattern)
	stringBackwards := fmt.Sprintf(`(\d|%s)`, reverseString(numbersPattern))
	regexpForwards := regexp.MustCompile(stringForwards)
	regexpBackwards := regexp.MustCompile(stringBackwards)

	first := regexpForwards.FindString(line)
	last := regexpBackwards.FindString(reverseString(line))

	number := stringToInt(first) * 10 + stringToInt(reverseString(last))
	log.Print("Part 2: ", number)
	return number
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func stringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err == nil {
		return number
	} else {
		value, err := strToNum[s]
		if !err {
			log.Panicf("%s is not a digit or a number", s)
		}
		return value
	}
}
