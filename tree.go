
package dtree

// A tree manage all the nodes.
type Tree struct {
	list []*Node      // The nodes of the current tree.
}

// Create a new tree from a dir,
// name is dir path.
func NewTree(name string) *Tree {
	t := &Tree{
		list: make([]*Node, 20),
	}

	return t
}

// Add a node to the tree.
func (t *Tree) AddNode(node *Node) {
	t.list = append(t.list, node)
}
