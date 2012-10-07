package i3ipc

import (
	"encoding/json"
	"net"
)

func GetMarks(ipc net.Conn) (marks []string, err error) {
	json_reply, err := Raw(ipc, I3GetMarks, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &marks)
	return
}
