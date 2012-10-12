package i3ipc

import (
	"encoding/json"
)


// Error for replies from a command to i3.
type CommandError string

func (self CommandError) Error() string {
	return string(self)
}

// Struct for replies from command messages.
type commandReply struct {
	Success bool
	Error   string
}

// Send a command to i3.
// FIXME: Doesn't support chained commands yet.
func Command(action string, ipc IPCSocket) (success bool, err error) {
	json_reply, err := ipc.Raw(I3Command, action)
	if err != nil {
		return
	}

	var cmd_reply []commandReply
	err = json.Unmarshal(json_reply, &cmd_reply)
	if err != nil {
		return
	}

	success = cmd_reply[0].Success
	if cmd_reply[0].Error == "" {
		err = nil
	} else {
		err = CommandError(cmd_reply[0].Error)
	}

	return
}
