
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
	list := make([]string, 0, 10)

	for _, node := range t.list {
		list = append(list, node.name)
	}

	output, _ := json.Marshal(list)
	return string(output)
}

// quick search
func (t *Tree) SearchNode(name string) *Node{
	list := t.list

	for _, item := range list {
		if item.GetName() == name {
			return item
		}
	}

	return nil
}
