
package scanner

import (
	_ "text/scanner"
)

func init() {
	AddScanner("js", &JsScanner{ScannerBase{}})
}

type JsScanner struct{
	ScannerBase
}

// Return the deps
func (s *JsScanner) Scan() []string {
	if !s.HasInit() {
		panic("Please init the scanner first.")
	}

	return nil
}

func (s *JsScanner) New() Scanner {
	return &JsScanner{ScannerBase{}}
}
