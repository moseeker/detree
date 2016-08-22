
package detree

import (
	"sync"
	"encoding/json"
)

// A tree manage all the nodes.
type Tree struct {
	lock *sync.Mutex
	// The nodes of the current tree.
	list []*Node  `json:"entry"`
}

// Create a new tree.
func NewTree() *Tree {
	t := &Tree{
		list: make([]*Node, 0, 20),
		lock: &sync.Mutex{},
	}

	return t
}

// Add a node to the tree.
// TODO: Check duplicate.
func (t *Tree) AddNode(node *Node) {
	t.list = append(t.list, node)
}

// Json output
func (t *Tree) JsonString() string {
	output := ""

	for _, node := range t.list {
		output += ("node: " + node.name + "\n")
	}
}
