
package main

import(
	"path"
	"github.com/towry/detree"
)

func main() {
	list := detree.ScanDir("./__test__")

	css := make([]string, 0, 10)
	js := make([]string, 0, 10)

	for _, item := range list {
		if path.Ext(item) == ".less" {
			css = append(css, item)
		} else if path.Ext(item) == ".js" {
			js = append(js, item)
		}
	}

	_ = detree.Build(css, nil)
	_ = detree.Build(js, nil)
}
