
package detree


// Context pass down to the sub modules.
type Context struct {
	root string     // The root dependant of the file.
	tree *Tree
}

// Create a new Context.
// name is the root path name.
func NewContext(name string, tree *Tree) *Context {
	return &Context{
		root: name,
		tree: tree,
	}
}

// Get the root.
func (c *Context) GetRoot() string {
	return c.root
}

// Set the root.
func (c *Context) SetRoot(root string) {
	c.root = root
}

// Get the tree
func (c *Context) GetTree() *Tree {
	return c.tree
}
