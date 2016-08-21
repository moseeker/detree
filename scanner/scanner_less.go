
package scanner

import (
	"io"
	"fmt"
	"github.com/tdewolff/parse/css"
)

func init() {
	AddScanner("less", &LessScanner{})
}

type LessScanner struct{
	lexer *css.Lexer
}

// Return the deps
func (s *LessScanner) Scan() []string {
	if s.lexer == nil {
		panic("Please init the scanner first.")
	}

	// deps := make([]string, 0, 10)

	for {
		tok, text := s.Next()

		switch tok {
		case css.ErrorToken:
			if s.lexer.Err() != io.EOF {
				fmt.Println("Error")
			}
			return
		case css.IdentToken:
			fmt.Println("Ident", string(text))
		case css.AtKeywordToken:
			fmt.Println("AtKeyword", string(text))
		}
	}

	return nil
}

// Create a new copy of the scanner.
func (s *LessScanner) New() Scanner {
	return &LessScanner{}
}

// Init the scanner with an io reader.
func (s *LessScanner) Init(in io.Reader) {
	s.lexer = css.NewLexer(in)
}

// Returns token and token text.
func (s *LessScanner) Next() (css.TokenType, []byte) {

	return s.lexer.Next()
}
