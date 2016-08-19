
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

// Return the deps
func (s *LessScanner) Scan() []string {
	if !s.HasInit() {
		panic("Please init the scanner first.")
	}

	return nil
}
