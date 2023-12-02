package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"strconv"
	)

func main() {
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
		numberPartOne := ExtractNumberPartOne(line)
		numberPartTwo := ExtractNumberPartTwo(line)
		accumPartOne = accumPartOne + numberPartOne
		accumPartTwo = accumPartTwo + numberPartTwo
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file", err)
	}


	fmt.Println("Result part one: ", accumPartOne)
	fmt.Println("Result part two: ", accumPartTwo)
}

func ExtractNumberPartOne(line string) int {
	r := regexp.MustCompile(`\d`)
	matches := r.FindAllString(line, -1)

	first := matches[0]
	last := matches[len(matches) - 1]
	numberString := first + last
	number, err := strconv.Atoi(numberString)

	if err != nil {
		log.Fatal("Error extracting number from ", numberString, err)
	}
	fmt.Println(number)
	return number
}


var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var strToNum = make(map[string]int)

func ExtractNumberPartTwo(line string) int {
	// Fill mapping of string to numbers
	for i, str := range numbers {
		strToNum[str] = i + 1
	}

	stringForwards := fmt.Sprintf(`(\d|%s)`, strings.Join(numbers, "|"))
	stringBackwards := fmt.Sprintf(`(\d|%s)`, strings.Join(mapReverse(numbers), "|"))
	regexpForwards := regexp.MustCompile(stringForwards)
	regexpBackwards := regexp.MustCompile(stringBackwards)

	first := regexpForwards.FindString(line)
	last := regexpBackwards.FindString(reverseString(line))

	number := stringToInt(first) * 10 + stringToInt(reverseString(last))
	fmt.Println(number)
	return number
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func mapReverse(arr []string) []string {
	result := make([]string, len(arr))

	for i, str := range arr {
		result[i] = reverseString(str)
	}

	return result
}

func stringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err == nil {
		return number
	} else {
		value, _ := strToNum[s]
		return value
	}
}
