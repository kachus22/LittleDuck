package main

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kachus22/LittleDuck/ANTLR/parser"
)

type TreeShapeListener struct {
	*parser.BaseLittleDuckListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

type SyntaxError struct {
	line, column int
	msg          string
}

func (s *SyntaxError) Error() string {
	return fmt.Sprintf("line %d:%d %s", s.line, s.column, s.msg)
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func NewErrorListener() *ErrorListener {
	return new(ErrorListener)
}

func (s *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	s.Errors = append(s.Errors, &SyntaxError{line, column, msg})
}

type TI struct {
	src   string
	valid bool
}

var testData = []*TI{
	{"../examples/good_input.txt", true},
	{"../examples/bad_input.txt", true},
}

func TestInputs(t *testing.T) {
	for _, ts := range testData {
		fmt.Printf("RUNNING\t%s\n", ts.src)
		// Setup the input
		input, _ := antlr.NewFileStream(ts.src)

		// Create the Lexer
		lexer := parser.NewLittleDuckLexer(input)

		// Load the Tokens
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		parser := parser.NewLittleDuckParser(stream)

		errorListener := NewErrorListener()
		parser.RemoveErrorListeners()
		parser.AddErrorListener(errorListener)

		parser.BuildParseTrees = true
		tree := parser.Programa()
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

		if ts.valid && len(errorListener.Errors) > 0 {
			for _, e := range errorListener.Errors {
				fmt.Printf("%s", ts.src)
				t.Error(e)
			}
		} else if ts.valid && len(errorListener.Errors) == 0 {
			fmt.Printf("SUCCESS\t%s\n\n", ts.src)
		}
	}
}
