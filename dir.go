
package detree

import (
	"os"
	"log"
	"path"
	"io/ioutil"
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
		list = append(list, path.Join(name, file.Name()))
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
		go BuildFile(list, ctx)

		// close the file
		reader.Close()
	}

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

// Build dependency for file
func BuildFile(list []string, ctx *Context) {
	// tree := ctx.GetTree()
	log.Println("build file", ctx.GetRoot())
}

func addFileToTree(name string, tree *Tree) {
	tree.lock.Lock()
	node := NewNode(name)
	tree.AddNode(node)
	tree.lock.Unlock()
}
