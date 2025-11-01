package repl

import (
	"bufio"
	"fmt"
	"io"
	"MagicInterpreter/lexer"
	"MagicInterpreter/parser"
	"MagicInterpreter/evaluator"
	"MagicInterpreter/object"
)

// PROMPT defines the text shown to the user each time
// the REPL is ready to accept input.
const PROMPT = ">>"


// Start launches the Read–Eval–Print Loop (REPL).
//
// It continuously:
//   1. Prints a prompt to the user
//   2. Reads a line of input from the provided io.Reader
//   3. Passes that input into the lexer to tokenize it
//   4. Prints out each token until the end of the line
//
// Parameters:
//   - in  : The input source (typically os.Stdin for keyboard input)
//   - out : The output destination (typically os.Stdout for console output)
//
// Example usage:
//
//   repl.Start(os.Stdin, os.Stdout)
//
// User types:  let x = 5;
// REPL prints: tokens for "let", "x", "=", "5", ";"
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// the E in REPL i.e evaluate/evaluator
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}


func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

