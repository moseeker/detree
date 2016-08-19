
package scanner

import (
	_ "text/scanner"
)

func init() {
	AddScanner("less", &LessScanner{ScannerBase{}})
}

type LessScanner struct{
	ScannerBase
}

// Return the next dep
func (s *LessScanner) Next() string {
	if !s.HasInit() {
		panic("Please init the scanner first.")
	}

	return ""
}
