
package dtree

import (
	"log"
	"path"
	"io/ioutil"
)

// Scan the given dir, get the file list.
func ScanDir(name string) []string {
	list := make([]string, 0, 10)
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		list = append(list, path.Join(name, file.Name()))
	}

	return list
}

