
package scanner

import "testing"

func TestScanner(t *testing.T) {
	_, err := GetScannerByName("js")
	if err != nil {
		t.Fatal()
	}
}

func TestScannerHasInit(t *testing.T) {
	s, _ := GetScannerByName("js")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not throw error")
		}
	}()

	// If the scanner has not init,
	// then call Next method will throw an error.
	s.Next()
}
