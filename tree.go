// Author: slowpoke <mail plus git at slowpoke dot io>
// Repository: https://github.com/proxypoke/i3ipc
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License.
// It comes without any warranty, to the extent permitted by
// applicable law. For a copy of the license, see COPYING or
// head to http://sam.zoy.org/wtfpl/COPYING.

package i3ipc

import (
	"encoding/json"
)

// Struct representing a Node in the i3 tree. For documentation of the fields,
// refer to http://i3wm.org/docs/ipc.html#_tree_reply.
type I3Node struct {
	Id                   int32
	Name                 string
	Type                 string
	Border               string
	Current_Border_Width int32
	Layout               string
	Orientation          string
	Percent              float64
	Rect                 Rect
	Window_Rect          Rect
	Deco_Rect            Rect
	Geometry             Rect
	Window               int32
	Urgent               bool
	Focused              bool
	Floating_Nodes       []I3Node
	Nodes                []I3Node
	Parent               *I3Node

	// Properties, not listed in docs:
	Window_Properties struct {
		// Transient_for ?
		Title    string
		Instance string
		Class    string
	}
	// Swallows []I3Node ?
	Sticky            bool
	Floating          string
	Last_Split_Layout string
	// Focus []I3Node ?
	Fullscreen_Mode  int32
	Scratchpad_State string
	Workspace_Layout string
}

// Traverses the tree setting correct reference to a parent node.
func setParent(node, parent *I3Node) {

	node.Parent = parent

	for i := range node.Nodes {
		setParent(&node.Nodes[i], node)
	}
	for i := range node.Floating_Nodes {
		setParent(&node.Floating_Nodes[i], node)
	}
}

// GetTree fetches the layout tree.
func (self *IPCSocket) GetTree() (root I3Node, err error) {
	json_reply, err := self.Raw(I3GetTree, "")
	if err != nil {
		return
	}

	defer setParent(&root, nil)

	err = json.Unmarshal(json_reply, &root)
	if err == nil {
		return
	}
	// For an explanation of this error silencing, see GetOutputs().
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		err = nil
	}
	return
}
