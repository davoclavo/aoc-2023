package main

import (
	p "github.com/ajitid/goparsify"
	"log"
	"reflect"
	"testing"
)

func TestColorParser(t *testing.T) {
	res, err := p.Run(colorParser, "1 red")
	log.Print(err)
	log.Print(res)
	log.Printf("%T", res)
	expected := Cubes{Count: 1, Color: "red"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal. \n%v\n%v", res, expected)
	}
}

func TestSetParser(t *testing.T) {
	res, err := p.Run(setParser, "1 red, 20 blue")
	log.Print(err)
	log.Print(res)
	log.Printf("%T", res)
	expected := []Cubes{
		{Count: 1, Color: "red"},
		{Count: 20, Color: "blue"},
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal. \nfound:    %v\nexpected: %v", res, expected)
	}
}

func TestSetsParser(t *testing.T) {
	res, err := p.Run(setsParser, "1 red, 2 blue; 2 red, 1 blue")
	log.Print(err)
	log.Print(res)
	log.Printf("%T", res)
	expected := []Set{
		[]Cubes{
			{Count: 1, Color: "red"},
			{Count: 2, Color: "blue"},
		},
		[]Cubes{
			{Count: 2, Color: "red"},
			{Count: 1, Color: "blue"},
		},
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal. \nfound:    %v\nexpected: %v", res, expected)
	}
}

func TestGameParser(t *testing.T) {
	res, err := p.Run(gameParser, "Game 10: 1 red, 2 blue; 2 red, 1 blue")
	log.Print(err)
	log.Print(res)
	log.Printf("%T", res)
	expected := Game{
		Id: 10,
		Sets: []Set{
			[]Cubes{
				{Count: 1, Color: "red"},
				{Count: 2, Color: "blue"},
			},
			[]Cubes{
				{Count: 2, Color: "red"},
				{Count: 1, Color: "blue"},
			},
		},
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal. \nfound:    %v\nexpected: %v", res, expected)
	}
}
