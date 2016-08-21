
package scanner

import (
	"io"
	"fmt"
	"github.com/tdewolff/parse/js"
)

func init() {
	AddScanner("js", &JsScanner{})
}

type JsScanner struct{
	lexer *js.Lexer
}

// Return the deps
func (s *JsScanner) Scan() []string {
	if s.lexer == nil {
		panic("Please init the scanner first.")
	}

	// deps := make([]string, 0, 10)

	for {
		tok, text := s.Next()

		switch tok {
		case js.ErrorToken:
			if s.Err() != io.EOF {
				fmt.Println("Error")
			}
			return nil
		case js.IdentifierToken:
			fmt.Println("Ident", string(text))
		}
	}

	return nil
}

// Create a new copy of the scanner.
func (s *JsScanner) New() Scanner {
	return &JsScanner{}
}

// Init the scanner with an io reader.
func (s *JsScanner) Init(in io.Reader) {
	s.lexer = js.NewLexer(in)
}

// Returns token and token text.
func (s *JsScanner) Next() (js.TokenType, []byte) {

	return s.lexer.Next()
}

// Check err
func (s *JsScanner) Err() error {
	return s.lexer.Err()
}
