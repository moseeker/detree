
package detree

// A tree manage all the nodes.
type Tree struct {
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
func (t *Tree) AddNode(node *Node) {
	t.list = append(t.list, node)
}
