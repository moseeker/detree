
package scanner

import (
	"io"
	"errors"
	"github.com/tdewolff/parse/css"
	"github.com/towry/detree"
)

func init() {
	AddScanner("less", &LessScanner{})
}

type LessScanner struct{
	lexer *css.Lexer
}

// Return the deps
func (s *LessScanner) Scan() ([]string, error) {
	if s.lexer == nil {
		panic("Please init the scanner first.")
	}

	deps := make([]string, 0, 10)

	for {
		tok, text := s.Next()

		switch tok {
		case css.ErrorToken:
			if s.lexer.Err() != io.EOF {
				return nil, s.Err()
			}
			return deps, nil
		case css.AtKeywordToken:
			keyword := string(text)
			if keyword == "@import" {
				dep, err := s.scanNextDep()
				if err != nil {
					break
				}
				deps = append(deps, dep)
			}
		}
	}
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

// Check err
func (s *LessScanner) Err() error {
	return s.lexer.Err()
}

// Scan the next dependency
func (s *LessScanner) scanNextDep() (string, error) {
	// expect string

	for {
		tok, text := s.Next()

		switch tok {
		case css.SemicolonToken:
		case css.ErrorToken:
		case css.BadStringToken:
			return "", errors.New("unexpected token")
		case css.StringToken:
			return detree.RemoveQuotes(string(text)), nil
		}
	}

	return "", errors.New("unexpected token")
}
