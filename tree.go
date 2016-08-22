
package detree

import (
	"sync"
)

// A tree manage all the nodes.
type Tree struct {
	lock *sync.Mutex
	list []*Node      // The nodes of the current tree.
}

// Create a new tree.
func NewTree() *Tree {
	t := &Tree{
		list: make([]*Node, 0, 20),
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
	return ""
}
