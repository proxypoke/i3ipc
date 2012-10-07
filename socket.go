package i3ipc

import (
	"bytes"
	"net"
	"os/exec"
	"strings"
	"unsafe"
)

// Magic string for the IPC API.
const MAGIC string = "i3-ipc"

// The types of messages that Raw() accepts.
type MessageType int32

const (
	I3Command MessageType = iota
	I3GetWorkspaces
	I3Subcribe
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
func Raw(ipc net.Conn, type_ MessageType, args string) (json_reply []byte, err error) {
	// Set up the parts of the message.
	var (
		message  []byte = []byte(MAGIC)
		payload  []byte = []byte(args)
		length   int32  = int32(len(payload))
		bytelen  [4]byte
		bytetype [4]byte
	)

	// Black Magic™.
	bytelen = *(*[4]byte)(unsafe.Pointer(&length))
	bytetype = *(*[4]byte)(unsafe.Pointer(&type_))

	for _, b := range bytelen {
		message = append(message, b)
	}
	for _, b := range bytetype {
		message = append(message, b)
	}
	for _, b := range payload {
		message = append(message, b)
	}

	_, err = ipc.Write(message)
	if err != nil {
		return
	}

	// Receive and assemble the reply.
	// Not sure if there's a cleaner solution but it seems to work.
	for {
		tmp := make([]byte, 1024)
		n, err := ipc.Read(tmp)

		for _, b := range tmp {
			json_reply = append(json_reply, b)
		}
		if n < 1024 || err != nil {
			break
		}
	}

	// Strip the first 14 bytes, which are the MAGIC string (6 bytes), the
	// payload length (4 bytes) and the message type (another 4 bytes).
	json_reply = json_reply[14:]
	// Get rid of trailing null bytes.
	json_reply = bytes.Trim(json_reply, "\x00")
	return
}
