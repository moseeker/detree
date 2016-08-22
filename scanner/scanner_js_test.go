
package scanner

import "os"
import "testing"


func TestScannerJs(t *testing.T) {
	reader, err := os.Open("../__test__/a.js")
	if err != nil {
		t.Fatal()
	}
	defer reader.Close()

	jss, _ := GetScannerByName("js")
	jss.Init(reader)

	list, _ := jss.Scan()

	if list == nil {
		t.Fatal()
	}
}
