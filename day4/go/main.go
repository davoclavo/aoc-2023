package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"

	p "github.com/ajitid/goparsify"
	mapset "github.com/deckarep/golang-set/v2"
)

func timer() func() {
	start := time.Now()
	return func() {
		log.Printf("Timer: %v", time.Since(start))
	}
}

func main() {
	defer timer()()
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows = []Row{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		log.Print("Reading line: ", line)

		res, err := p.Run(rowParser, line)
		log.Print("err ", err)
		log.Printf("res: %+v", res)
		log.Printf("res type: %t", res)
		row := res.(Row)
		if err != nil {
			log.Panicf("Error parsing line: \n%s", err)
			panic("Error parsing line")
		}
		rows = append(rows, row)
	}

	numberPartOne := calculateNumberPartOne(rows)
	numberPartTwo := calculateNumberPartTwo(rows)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file", err)
	}

	log.Print("----------------")
	log.Print("Result part one: ", numberPartOne)
	log.Print("Result part two: ", numberPartTwo)
}

var IntLit = p.Regex("[0-9]+").Map(func(n *p.Result) {
	num, _ := strconv.Atoi(n.Token)
	n.Result = num
})

type Row struct {
	WinningNumbers mapset.Set[int]
  Numbers []int
}

var rowParser = p.Seq(
	"Card",
	IntLit,
	":",
	p.OneOrMore(
		IntLit,
	),
	"|",
	p.OneOrMore(
		IntLit,
	),
).Map(func(n *p.Result) {
	var winningNumbers = mapset.NewSet[int]()
	var numbers = []int{}

	for _, child := range n.Child[3].Child {
		winningNumbers.Add(child.Result.(int))
	}
	for _, child := range n.Child[5].Child {
		numbers = append(numbers, child.Result.(int))
	}
	n.Result = Row{WinningNumbers: winningNumbers, Numbers: numbers}
})

func calculateNumberPartOne(rows []Row) int {
	var accumulator int = 0

	for _, row := range rows {
		var rowAccumulator = 1

		for _, number := range row.Numbers {
			if row.WinningNumbers.Contains(number) {
				rowAccumulator = rowAccumulator << 1
			}
		}

		rowAccumulator = rowAccumulator >> 1

		accumulator = accumulator + rowAccumulator
	}


	return accumulator
}


func calculateNumberPartTwo(rows []Row) int {
	// cardCounts is an slice of the same size as rows filled with ones
	cardCounts := make([]int, len(rows))

	for i, row := range rows {
		matchingNumbers := 0
		for _, number := range row.Numbers {
			if row.WinningNumbers.Contains(number) {
				matchingNumbers++
				cardCounts[i+matchingNumbers] += cardCounts[i] + 1
			}
		}
	}

	accumulator := len(cardCounts)
	for _, count := range cardCounts {
		accumulator += count
	}

	return accumulator
}


