
package scanner 

import (
	"io"
	"fmt"
)

// Scanner interface.
type Scanner interface {
	Scan() []string 	 // Get the next dep the file requires.
	Init(io.Reader)      // User must init the scanner with a reader.
	// Internal api.
	New() Scanner        // Create a new scanner.
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
