// Package emailaddress_test contains tests for the emailaddress grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by maketest.go
//
package emailaddress_test

import (
	"bramp.net/antlr4-grammars/emailaddress"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/emailaddress/examples/example1.txt",
	"grammars-v4/emailaddress/examples/example2.txt",
	"grammars-v4/emailaddress/examples/example3.txt",
	"grammars-v4/emailaddress/examples/example4.txt",
	"grammars-v4/emailaddress/examples/example5.txt",
	"grammars-v4/emailaddress/examples/example6.txt",
	"grammars-v4/emailaddress/examples/example7.txt",
	"grammars-v4/emailaddress/examples/example8.txt",
	"grammars-v4/emailaddress/examples/example9.txt",
}

type exampleListener struct {
	*emailaddress.BaseemailaddressListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Example() {
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")

	// Create the Lexer
	lexer := emailaddress.NewemailaddressLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := emailaddress.NewemailaddressParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Emailaddress()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)
}

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filepath.Join("..", filename))
	if err != nil {
		return nil, err
	}

	return input, nil
}

func TestemailaddressLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := emailaddress.NewemailaddressLexer(input)

		// Try and read all tokens
		i := 0
		for ; i < MAX_TOKENS; i++ {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		// If we read too many tokens, then perhaps there is a problem with the lexer.
		if i == MAX_TOKENS {
			t.Errorf("NewemailaddressLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestemailaddressParser(t *testing.T) {
	// TODO(bramp): Run this test with and without p.BuildParseTrees

	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := emailaddress.NewemailaddressLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := emailaddress.NewemailaddressParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true)) // TODO Change this
		p.AddErrorListener(antlr.NewConsoleErrorListener())

		// Finally test
		p.Emailaddress()

		// TODO Check for errors
	}
}
