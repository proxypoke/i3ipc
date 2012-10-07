package i3ipc

import (
	"bytes"
	"encoding/json"
	"net"
	"os/exec"
	"strings"
)

// The types of messages that Raw() accepts.
type MessageType int

const (
	I3Command MessageType = iota
	I3GetWorkspaces
	I3GetOutputs
	I3GetTree
	I3GetMarks
	I3GetBarConfig
	I3GetVersion
)

// Error for unknown message types.
type MessageTypeError string

func (self MessageTypeError) Error() string {
	return string(self)
}

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
func Raw(type_ MessageType, args string, _ net.Conn) (json_reply []byte, err error) {
	var (
		out        bytes.Buffer
		typestring string
	)

	switch type_ {
	case I3Command:
		typestring = "command"
	case I3GetWorkspaces:
		typestring = "get_workspaces"
	case I3GetOutputs:
		typestring = "get_outputs"
	case I3GetTree:
		typestring = "get_tree"
	case I3GetMarks:
		typestring = "get_marks"
	case I3GetBarConfig:
		typestring = "get_bar_config"
	case I3GetVersion:
		typestring = "get_version"
	default:
		err = MessageTypeError("Unknown message type.")
		return
	}

	cmd := exec.Command("i3-msg", "-t", typestring, args)
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
