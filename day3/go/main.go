package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"time"
	"slices"

	p "github.com/ajitid/goparsify"
)

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
	var rows = []Row{}
	for scanner.Scan() {
		line := scanner.Text()
		log.Print("Reading line: ", line)

		res, err := p.Run(rowParser, line)
		// log.Print(err)
		// log.Print(res)
		// log.Printf("%t", res)
		row := res.(Row)
		if err != nil {
			log.Panicf("Error parsing line: \n%s", err)
			panic("Error parsing line")
		}
		rows = append(rows, row)
	}

	numberPartOne := calculateNumberPartOne(rows)

	// numberPartTwo := extractNumberPartTwo(game)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file", err)
	}

	log.Print("----------------")
	log.Print("Result part one: ", numberPartOne)
	// log.Print("Result part two: ", numberPartTwo)
}

type Number struct {
	X int64
	Value int64
}

type Symbol struct {
	X int64
}

type Row struct {
  Numbers []Number
  Symbols []Symbol
}

var IntLit = p.Regex("[0-9]+").Map(func(n *p.Result) {
	num, _ := strconv.Atoi(n.Token)
	n.Result = num
})

var rowParser = p.Seq(
	p.ZeroOrMore("."),
	p.OneOrMore(
		p.Any(IntLit, p.NotChars(".", 0, 1)),
		p.ZeroOrMore("."),
	),
).Map(func(n *p.Result) {
	var numbers = []Number{}
	var symbols = []Symbol{}

	for _, child := range n.Child[1].Child {
		if child.Result == nil {
			symbols = append(symbols, Symbol{X: int64(child.Start)})
		} else {
			switch child.Result.(type) {
			// NOTE: goparsify parses 134. as float64 and 134 as int64
			case float64:
				numbers = append(numbers, Number{X: int64(child.Start), Value: int64(child.Result.(float64))})
			case int64:
				numbers = append(numbers, Number{X: int64(child.Start), Value: child.Result.(int64)})
			case int:
				numbers = append(numbers, Number{X: int64(child.Start), Value: int64(child.Result.(int))})
		  default:
				panic("Invalid type for number!")
			}
		}
	}
	n.Result = Row{Numbers: numbers, Symbols: symbols}
})

func calculateNumberPartOne(rows []Row) int {

	var accumulator = 0

	for idx, row := range rows {
		for _, number := range row.Numbers {
			minRow := int(math.Max(0, float64(idx) - 1))
			maxRow := int(math.Min(float64(len(rows)), float64(idx) + 1))
			minCol := number.X - 1
			maxCol := number.X + int64(math.Trunc(math.Log10(float64(number.Value)))) + 1

			for _, applicableRow := range rows[minRow : maxRow + 1] {
				symbols := applicableRow.Symbols
				symbolIdx := slices.IndexFunc(symbols, func(s Symbol) bool { return s.X >= minCol && s.X <= maxCol })
				if symbolIdx != -1 {
					accumulator = accumulator + int(number.Value)
					// log.Printf("found! %v, %v - %v", number.X, idx, number.Value)
					// log.Printf("range: %v, %v - %v, %v", minRow, maxRow, minCol, maxCol)
					// log.Printf("symbol! %v, %v", symbols, symbolIdx)
					break
				}
			}

		}

	}
	return accumulator
}
