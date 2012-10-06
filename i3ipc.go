package i3ipc

import (
	"bytes"
	"encoding/json"
	"net"
	"os/exec"
	"strings"
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

// Connect to the i3 IPC socket and return it.
func GetIPCSocket() (ipc net.Conn, err error) {
	var out bytes.Buffer

	cmd := exec.Command("i3", "--get-socketpath")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}

	path := strings.TrimSpace(out.String())
	ipc, err = net.Dial("unix", path)
	return
}

// Send raw messages to i3. Returns a json bytestring.
// FIXME: Uses exec to access i3-msg for now. Should use the socket instead.
// FIXME: type_ should be an enumerated constant, not a string.
func Raw(type_, args string, _ net.Conn) (json_reply []byte, err error) {
	var out bytes.Buffer

	cmd := exec.Command("i3-msg", "-t", type_, args)
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}

	json_reply = out.Bytes()
	return
}

// Send a command to i3.
// FIXME: Doesn't support chained commands yet.
func Command(action string, ipc net.Conn) (success bool, err error) {
	json_reply, err := Raw("command", action, ipc)
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
