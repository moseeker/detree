
package detree

import (
	"fmt"
	"testing"
)

func TestScanDir(t *testing.T) {
	dir := "./__test__"
	list := ScanDir(dir)

	if len(list) <= 0 {
		t.Fatal()
	}

	fmt.Printf("%q", list)
}
