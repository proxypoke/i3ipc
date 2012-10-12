package i3ipc

import (
	"encoding/json"
)

type I3Node struct {
	Id                   int32
	Name                 string
	Border               string
	Current_Border_Width int32
	Layout               string
	Percent              float64
	Rect                 Rect
	Window_Rect          Rect
	Geometry             Rect
	Window               int32
	Urgent               bool
	Focused              bool
	Nodes                []I3Node
}

func GetTree(ipc IPCSocket) (root I3Node, err error) {
	json_reply, err := ipc.Raw(I3GetTree, "")
	if err != nil {
		return
	}

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
