
package detree

import (
	"os"
	"fmt"
	"log"
	"time"
	"path"
	"io/ioutil"
	"path/filepath"
	"github.com/towry/detree/scanner"
)

// Scan the given dir, get the file list.
func ScanDir(name string) []string {
	list := make([]string, 0, 10)
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := path.Join(name, file.Name())
		filename, err = filepath.Abs(filename)
		if err != nil {
			continue
		}

		list = append(list, filename)
	}

	return list
}


// files is the entry files.
// config is build config.
func Build(files []string, config *Config) *Tree {
	tree := NewTree()

	for _, file := range files {
		scannerName := getScannerFromPath(file)
		// build the dependency tree for the file
		scanner, err := scanner.GetScannerByName(scannerName)
		if err != nil {
			continue
		}

		reader, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		// init the scanner.
		scanner.Init(reader)
		list, errScan := scanner.Scan()
		if errScan != nil {
			log.Fatal(errScan)
		}

		addFileToTree(file, tree)

		ctx := NewContext(file, tree)
		
		go BuildEachFile(list, ctx)

		// close the file
		reader.Close()
	}

	time.Sleep(time.Millisecond)

	return tree
}

// Get scanner from path extension.
func getScannerFromPath(name string) string {
	ext := path.Ext(name)

	if ext == "" {
		return ""
	}

	return ext[1:]
}

// Build dependency for each file
// list is the dependencies.
func BuildEachFile(list []string, ctx *Context) {
	// tree := ctx.GetTree()

	for _, file := range list {
		fmt.Println(file)
	}
}

func addFileToTree(name string, tree *Tree) {
	tree.lock.Lock()
	found := tree.SearchNode(name)
	// if added, return
	if found != nil {
		return
	}

	node := NewNode(name)
	tree.AddNode(node)
	tree.lock.Unlock()
}
