package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if tok.Type == token.ILLEGAL {
				// TODO: Error handling
				fmt.Println("Parse Error.")
				break
			}
			fmt.Printf("%+v\n", tok)
		}
	}
}
