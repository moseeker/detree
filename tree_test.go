
package detree

import (
	"testing"
)

func TestCreateTree(t *testing.T) {
	tree := NewTree()
	_ = tree
}

func TestAddNodeToTree(t *testing.T) {
	tree := NewTree()
	node := NewNode("./__test__")
	tree.AddNode(node)
}
