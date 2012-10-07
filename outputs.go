package i3ipc

import (
	"encoding/json"
	"net"
)

type Output struct {
	Name              string
	Active            bool
	Rect              Rect
	Current_Workspace string
	//Primary           bool
}

func GetOutputs(ipc net.Conn) (outputs []Output, err error) {
	json_reply, err := Raw(ipc, I3GetOutputs, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &outputs)
	if err == nil {
		return
	}
	// Outputs which aren't displaying any workspace will have JSON-null set as
	// their value for current_workspace. Since Go's equivalent, nil, can't be
	// assigned to strings, it will cause Unmarshall to return with an
	// UnmarshalTypeError, but otherwise correctly unmarshal the JSON input. We
	// simply ignore this error due to this reason.
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		err = nil
	}
	return
}
