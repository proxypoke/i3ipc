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

// Workspace represents a workspace. For documentation of the fields,
// refer to http://i3wm.org/docs/ipc.html#_workspaces_reply.
type Workspace struct {
	Num     int32
	Name    string
	Visible bool
	Focused bool
	Rect    Rect
	Output  string
	Urgent  bool
}

// Rect represents the geometry of a window, output or workspace.
type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// GetWorkspaces fetches a list of all current workspaces.
func (socket *IPCSocket) GetWorkspaces() (workspaces []Workspace, err error) {
	jsonReply, err := socket.Raw(I3GetWorkspaces, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &workspaces)
	return
}
