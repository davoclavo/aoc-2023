package main

import (
	p "github.com/ajitid/goparsify"
	"log"
	"reflect"
	"testing"
)

func TestRowParser(t *testing.T) {
	res, err := p.Run(rowParser, ".$467.114*-.")
	log.Print(err)
	log.Print(res)
	log.Printf("%t", res)
	expected := Row{
		Numbers: []Number{
			{X: 2, Value: 467},
			{X: 6, Value: 114},
		},
		Symbols: []Symbol{
			{X: 1},
			{X: 9},
			{X: 10},
		},
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal. \n%v\n%v", res, expected)
	}
}
