package i3ipc

import (
	"regexp"
	"strings"
)

// Returns the root node of the tree.
func (self *I3Node) Root() *I3Node {
	if self.Parent == nil {
		return self
	}
	return self.Parent.Root()
}

// Returns a slice of all descendent nodes.
func (self *I3Node) Descendents() []*I3Node {

	var collectDescendents func(*I3Node, []*I3Node) []*I3Node

	// Collects descendent nodes recursively
	collectDescendents = func(n *I3Node, collected []*I3Node) []*I3Node {
		for i := range n.Nodes {
			collected = append(collected, &n.Nodes[i])
			collected = collectDescendents(&n.Nodes[i], collected)
		}
		for i := range n.Floating_Nodes {
			collected = append(collected, &n.Floating_Nodes[i])
			collected = collectDescendents(&n.Floating_Nodes[i], collected)
		}
		return collected
	}

	return collectDescendents(self, nil)
}

// Returns nodes that has no children nodes (leaves).
func (self *I3Node) Leaves() (leaves []*I3Node) {

	nodes := self.Descendents()

	for i := range nodes {
		node := nodes[i]

		is_dockarea := node.Parent != nil && node.Parent.Type == "dockarea"

		if len(node.Nodes) == 0 && node.Type == "con" && !is_dockarea {
			leaves = append(leaves, node)
		}
	}
	return
}

// Returns all nodes of workspace type.
func (self *I3Node) Workspaces() (workspaces []*I3Node) {

	nodes := self.Descendents()

	for i := range nodes {
		i3_special := strings.HasPrefix(nodes[i].Name, "__")
		if nodes[i].Type == "workspace" && !i3_special {
			workspaces = append(workspaces, nodes[i])
		}
	}
	return
}

// Returns a node that is being focused now.
func (self *I3Node) FindFocused() *I3Node {

	nodes := self.Descendents()

	for i := range nodes {
		if nodes[i].Focused {
			return nodes[i]
		}
	}
	return nil
}

// Returns a node that has given id.
func (self *I3Node) FindByID(id int32) *I3Node {

	nodes := self.Descendents()

	for i := range nodes {
		if nodes[i].Id == id {
			return nodes[i]
		}
	}
	return nil
}

// Returns a node that has given window id.
func (self *I3Node) FindByWindow(window int32) *I3Node {

	nodes := self.Descendents()

	for i := range nodes {
		if nodes[i].Window == window {
			return nodes[i]
		}
	}
	return nil
}

// Returns nodes which name matches the regexp.
func (self *I3Node) FindNamed(name string) []*I3Node {

	nodes := self.Descendents()
	reName := regexp.MustCompile(name)
	var found []*I3Node

	for _, node := range nodes {
		if reName.MatchString(node.Name) {
			found = append(found, node)
		}
	}
	return found
}

// Looks for a workspace up the tree.
func (self *I3Node) Workspace() *I3Node {

	if self.Parent == nil { // the root node
		return nil
	}
	if self.Parent.Type == "workspace" {
		return self.Parent
	}

	return self.Parent.Workspace()
}
