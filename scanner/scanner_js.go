
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

// Return the next dep
func (s *JsScanner) Next() string {
	if !s.HasInit() {
		panic("Please init the scanner first.")
	}

	return ""
}
