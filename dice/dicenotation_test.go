// Package dice_test contains tests for the DiceNotation grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by maketest.go
//
package dice_test

import (
	"bramp.net/antlr4-grammars/dice"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/dice/examples/big_d.txt",
	"grammars-v4/dice/examples/big_dice.txt",
	"grammars-v4/dice/examples/dice_arithmetic.txt",
	"grammars-v4/dice/examples/dice_arithmetic.txt.tree",
	"grammars-v4/dice/examples/dice_with_sub.txt",
	"grammars-v4/dice/examples/dice_with_sum.txt",
	"grammars-v4/dice/examples/invalid_space_end.txt",
	"grammars-v4/dice/examples/invalid_space_end.txt.errors",
	"grammars-v4/dice/examples/invalid_space_start.txt",
	"grammars-v4/dice/examples/invalid_space_start.txt.errors",
	"grammars-v4/dice/examples/missing_dice.txt",
	"grammars-v4/dice/examples/missing_dice.txt.errors",
	"grammars-v4/dice/examples/missing_side.txt",
	"grammars-v4/dice/examples/missing_side.txt.errors",
	"grammars-v4/dice/examples/mixed_arithmetic.txt",
	"grammars-v4/dice/examples/mixed_arithmetic.txt.tree",
	"grammars-v4/dice/examples/negative_dice.txt",
	"grammars-v4/dice/examples/negative_dice.txt.errors",
	"grammars-v4/dice/examples/negative_number.txt",
	"grammars-v4/dice/examples/negative_number.txt.errors",
	"grammars-v4/dice/examples/no_dice.txt",
	"grammars-v4/dice/examples/num_arithmetic.txt",
	"grammars-v4/dice/examples/number.txt",
	"grammars-v4/dice/examples/simple_dice.txt",
}

type exampleListener struct {
	*dice.BaseDiceNotationListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Example() {
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")

	// Create the Lexer
	lexer := dice.NewDiceNotationLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := dice.NewDiceNotationParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Parse()
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

func TestDiceNotationLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := dice.NewDiceNotationLexer(input)

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
			t.Errorf("NewDiceNotationLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestDiceNotationParser(t *testing.T) {
	// TODO(bramp): Run this test with and without p.BuildParseTrees

	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := dice.NewDiceNotationLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := dice.NewDiceNotationParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true)) // TODO Change this
		p.AddErrorListener(antlr.NewConsoleErrorListener())

		// Finally test
		p.Parse()

		// TODO Check for errors
	}
}
