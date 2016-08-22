
package scanner

import "os"
import "testing"
import "fmt"


func TestScannerLess(t *testing.T) {
	reader, err := os.Open("../__test__/c.less")
	if err != nil {
		t.Fatal()
	}
	defer reader.Close()

	less, _ := GetScannerByName("less")
	less.Init(reader)

	list, _ := less.Scan()
	fmt.Printf("less: %q\n", list)

	if list == nil {
		t.Fatal()
	}
}
