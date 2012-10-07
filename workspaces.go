package i3ipc

import (
	"encoding/json"
	"net"
)

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Workspace struct {
	Num     int
	Name    string
	Visible bool
	Focused bool
	Rect    Rect
	Output  string
	Urgent  bool
}

func GetWorkspaces(ipc net.Conn) (workspaces []Workspace, err error) {
	json_reply, err := Raw(I3GetWorkspaces, "", ipc)
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &workspaces)
	return
}
