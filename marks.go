package i3ipc

import (
	"encoding/json"
)

func (self *IPCSocket) GetMarks() (marks []string, err error) {
	json_reply, err := self.Raw(I3GetMarks, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &marks)
	return
}
