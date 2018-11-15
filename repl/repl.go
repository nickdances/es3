package repl

import (
	"bufio"
	"fmt"	
	"io"
	"os"

	"es3/lexer"
	"es3/token"
)
//Start reads lines in and writes tokens out
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("\nES3 $ " + os.Getenv("USER") + " ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for lexed := lex.NextToken(); lexed.Type != token.EOF; lexed = lex.NextToken() {
			fmt.Printf("%+v\n", lexed)
		}
	}
}
