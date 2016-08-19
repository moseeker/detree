
package dtree

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
		deps: make([]string, 10),
		refs: make([]string, 10),
	}

	return n
}
