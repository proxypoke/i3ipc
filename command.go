package i3ipc

import (
	"encoding/json"
	"net"
)

// Struct for replies from command messages.
type CommandReply struct {
	Success bool
	Error   string
}

// Error for replies from a command to i3.
type CommandError string

func (self CommandError) Error() string {
	return string(self)
}

// Send a command to i3.
// FIXME: Doesn't support chained commands yet.
func Command(action string, ipc net.Conn) (success bool, err error) {
	json_reply, err := Raw(I3Command, action, ipc)
	if err != nil {
		return
	}

	var cmd_reply []CommandReply
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
