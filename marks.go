package i3ipc

import (
	"encoding/json"
)

func GetMarks(ipc IPCSocket) (marks []string, err error) {
	json_reply, err := ipc.Raw(I3GetMarks, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &marks)
	return
}
