package main

import (
	"bufio"
	p "github.com/ajitid/goparsify"
	"log"
	"os"
	"time"
)

var maxCubes = map[string]int64{
	"red":   12,
	"green": 13,
	"blue":  14,
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

		res, err := p.Run(gameParser, line)
		game := res.(Game)
		log.Print(game)
		log.Print(err)
		if err != nil {
			log.Panicf("Error parsing line: \n%s", err)
		}
		numberPartOne := extractNumberPartOne(game)
		numberPartTwo := extractNumberPartTwo(game)
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

type Cubes struct {
	Count int64
	Color string
}

type Set = []Cubes

type Game struct {
	Id   int64
	Sets []Set
}

var colorParser = p.Seq(
	p.NumberLit(),
	p.Any("red", "green", "blue"),
).Map(func(n *p.Result) {
	count := n.Child[0].Result.(int64)
	// log.Print(n.Child[1].Result)
	// log.Print(n.Child[1].Token)
	color := n.Child[1].Token
	n.Result = Cubes{Count: count, Color: color}
})

var setParser = p.OneOrMore(
	colorParser,
	",",
).Map(func(n *p.Result) {
	var set = Set{}
	for _, cubes := range n.Child {
		set = append(set, cubes.Result.(Cubes))
	}
	n.Result = set
})

var setsParser = p.OneOrMore(
	setParser,
	p.Maybe(";"),
).Map(func(n *p.Result) {
	var sets = []Set{}
	for _, set := range n.Child {
		sets = append(sets, set.Result.(Set))
	}
	n.Result = sets
})

var gameParser = p.Seq(
	"Game", p.NumberLit(), ":", setsParser,
).Map(func(n *p.Result) {
	n.Result = Game{Id: n.Child[1].Result.(int64), Sets: n.Child[3].Result.([]Set)}
})

func extractNumberPartOne(game Game) int {
	var number = game.Id
	for _, set := range game.Sets {
		for _, cubes := range set {
			maxCount, _ := maxCubes[cubes.Color]
			if cubes.Count > maxCount {
				number = 0
			}
		}
	}
	log.Print("Part 1: ", number)
	return int(number)
}

func extractNumberPartTwo(game Game) int {

	var minCubes = map[string]int64{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range game.Sets {
    for _, cube := range set {
			minColor, _ := minCubes[cube.Color]
			if cube.Count > minColor {
				minCubes[cube.Color] = cube.Count
			}
		}
	}

	minRed, _ := minCubes["red"]
	minGreen, _ := minCubes["green"]
	minBlue, _ := minCubes["blue"]

	number := int(minRed * minGreen * minBlue)
	log.Print("Part 2: ", number)
	return number
}
