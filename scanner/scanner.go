
package scanner 

import (
	"io"
	"fmt"
	"text/scanner"
)

// Scanner interface.
type Scanner interface {
	Scan() []string 	// Get the next dep the file requires.
	Init(io.Reader)     // Init the scanner
	New() Scanner
}

type ScannerBase struct {
	s *scanner.Scanner
}

func (s *ScannerBase) Scan() []string {
	return nil
}

func (s *ScannerBase) New() Scanner {
	return nil
}

func (s *ScannerBase) Init(in io.Reader) {
	s.s = new(scanner.Scanner).Init(in)
}

func (s *ScannerBase) HasInit() bool {
	if s.s == nil {
		return false
	}
	return true
}

// The supported scanners.
var Scanners map[string]Scanner

func init() {
	Scanners = make(map[string]Scanner)
}

// Add a scanner to the scanners.
// the scanner name should be the extension of 
// the file you want to scan.
// For example, 'js' for js files.
func AddScanner(name string, scanner Scanner) {
	Scanners[name] = scanner
}

// Get scanner
func GetScannerByName(name string) (Scanner, error) {
	scanner, ok := Scanners[name]

	if !ok {
		return nil, fmt.Errorf("No such scanner: %q\n", name)
	}

	return scanner.New(), nil
}
