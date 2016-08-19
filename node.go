
package detree

// A node represents file, has reference to
// the dependencies and dependants
type Node struct {
	name string       // The file full name
	deps []string     // The dependencies of file name
	refs []string     // The top most dependants
}

// Create a new node.
// name is the file's path
func NewNode(name string) *Node {
	n := &Node{
		name: name,
		deps: make([]string, 0, 10),
		refs: make([]string, 0, 10),
	}

	return n
}

// Add a dep to this node.
// !Is this operation of append item to slice
// efficient.
func (n *Node) AddDep(name string) {
	if n.HasDep(name) {
		return
	}

	n.deps = append(n.deps, name)
}

// Add a ref to this node.
func (n *Node) AddRef(name string) {
	if n.HasRef(name) {
		return
	}

	n.refs = append(n.refs, name)
}

// Test if the node has this dep.
func (n *Node) HasDep(name string) bool {
	deps := n.deps

	for _, dep := range deps {
		if dep == name {
			return true
		}
	}

	return false
}

// Test if the node has this ref.
func (n *Node) HasRef(name string) bool {
	refs := n.refs;

	for _, ref := range refs {
		if ref == name {
			return true
		}
	}

	return false
}
