// Package abnf_test contains tests for the Abnf grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by maketest.go
//
package abnf_test

import (
	"bramp.net/antlr4-grammars/abnf"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/abnf/examples/iri.abnf",
	"grammars-v4/abnf/examples/postal.abnf",
	"grammars-v4/abnf/examples/rfc5322.abnf",
}

type exampleListener struct {
	*abnf.BaseAbnfListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Example() {
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")

	// Create the Lexer
	lexer := abnf.NewAbnfLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := abnf.NewAbnfParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Rulelist()
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

func TestAbnfLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := abnf.NewAbnfLexer(input)

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
			t.Errorf("NewAbnfLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestAbnfParser(t *testing.T) {
	// TODO(bramp): Run this test with and without p.BuildParseTrees

	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := abnf.NewAbnfLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := abnf.NewAbnfParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true)) // TODO Change this
		p.AddErrorListener(antlr.NewConsoleErrorListener())

		// Finally test
		p.Rulelist()

		// TODO Check for errors
	}
}
