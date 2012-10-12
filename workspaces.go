package i3ipc

import (
	"encoding/json"
)

type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type Workspace struct {
	Num     int32
	Name    string
	Visible bool
	Focused bool
	Rect    Rect
	Output  string
	Urgent  bool
}

func GetWorkspaces(ipc IPCSocket) (workspaces []Workspace, err error) {
	json_reply, err := ipc.Raw(I3GetWorkspaces, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &workspaces)
	return
}
