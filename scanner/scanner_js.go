
package scanner

import (
	"io"
	"errors"
	"github.com/tdewolff/parse/js"
)

func init() {
	AddScanner("js", &JsScanner{})
}

type JsScanner struct{
	lexer *js.Lexer
}

// Return the deps
func (s *JsScanner) Scan() ([]string, error) {
	if s.lexer == nil {
		panic("Please init the scanner first.")
	}

	deps := make([]string, 0, 10)

	for {
		tok, text := s.Next()

		switch tok {
		case js.ErrorToken:
			if s.Err() != io.EOF {
				return nil, s.Err()
			}
			return deps, nil
		case js.IdentifierToken:
			keyword := string(text)
			if keyword == "require" || keyword == "import" {
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

// Scan the dependency
func (s *JsScanner) scanNextDep() (string, error) {
	// expect string

	for {
		tok, text := s.Next()

		switch tok {
		case js.LineTerminatorToken:
		case js.ErrorToken:
			return "", errors.New("unexpected token")
		case js.StringToken:
			return RemoveQuotes(string(text)), nil
		}
	}

	return "", errors.New("unexpected token")
}

