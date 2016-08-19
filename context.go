
package dtree


// Context pass down to the sub modules.
type Context struct {
	root string     // The root dependant of the file.
}

// Create a new Context.
// name is the root path name.
func NewContext(name string) *Context {
	return &Context{
		root: name,
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
