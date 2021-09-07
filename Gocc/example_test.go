package main

import (
	"fmt"
	"testing"

	"github.com/kachus22/LittleDuck/Gocc/lexer"
	"github.com/kachus22/LittleDuck/Gocc/parser"
)

type TI struct {
	src   string
	valid bool
}

var testData = []*TI{
	{"../examples/good_input.txt", true},
	{"../examples/bad_input.txt", true},
}

func TestInputs(t *testing.T) {
	p := parser.NewParser()
	for _, ts := range testData {
		fmt.Printf("RUNNING\t%s\n", ts.src)
		l, _ := lexer.NewLexerFile(ts.src)
		_, err := p.Parse(l)
		if ts.valid && err != nil {
			t.Error(err)
		} else {
			fmt.Printf("SUCCESS\t%s\n\n", ts.src)
		}
	}
}
