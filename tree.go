// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i3ipc

import (
	"encoding/json"
)

// I3Node represents a Node in the i3 tree. For documentation of the fields,
// refer to http://i3wm.org/docs/ipc.html#_tree_reply.
type I3Node struct {
	ID                 int32
	Name               string
	Border             string
	CurrentBorderWidth int32
	Layout             string
	Percent            float64
	Rect               Rect
	WindowRect         Rect
	Geometry           Rect
	Window             int32
	Urgent             bool
	Focused            bool
	Nodes              []I3Node
}

// GetTree fetches the layout tree.
func (socket *IPCSocket) GetTree() (root I3Node, err error) {
	jsonReply, err := socket.Raw(I3GetTree, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &root)
	if err == nil {
		return
	}
	// For an explanation of this error silencing, see GetOutputs().
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		err = nil
	}
	return
}
