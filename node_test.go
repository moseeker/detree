
package detree

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	node := NewNode("./__test__")
	_ = node
}

func TestAddDep(t *testing.T) {
	node := NewNode("./__test__")
	node.AddDep("../a.js")

	if !node.HasDep("../a.js") {
		t.Fatal()
	}
}

func TestAddRef(t *testing.T) {
	node := NewNode("./__test__")
	node.AddRef("../a.js")

	if !node.HasRef("../a.js") {
		t.Fatal()
	}
}

