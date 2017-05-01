// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i3ipc

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"unsafe"
)

const (
	// Magic string for the IPC API.
	_Magic string = "i3-ipc"
	// The length of the i3 message "header" is 14 bytes: 6 for the _Magic
	// string, 4 for the length of the payload (int32 in native byte order) and
	// another 4 for the message type (also int32 in NBO).
	_Headerlen = 14
)

// A Message from i3. Can either be a Reply or an Event.
type Message struct {
	Payload []byte
	IsEvent bool
	Type    int32
}

// A MessageType that Raw() accepts.
type MessageType int32

const (
	I3Command MessageType = iota
	I3GetWorkspaces
	I3Subscribe
	I3GetOutputs
	I3GetTree
	I3GetMarks
	I3GetBarConfig
	I3GetVersion
)

// MessageTypeError for unknown message types.
type MessageTypeError string

func (error MessageTypeError) Error() string {
	return string(error)
}

// MessageError for communication failures.
type MessageError string

func (error MessageError) Error() string {
	return string(error)
}

// IPCSocket represents a Unix socket to communicate with i3.
type IPCSocket struct {
	socket      net.Conn
	open        bool
	subscribers []chan Event
}

// Close the connection to the underlying Unix socket.
func (socket *IPCSocket) Close() error {
	socket.open = false
	return socket.socket.Close()
}

// GetIPCSocket creates a new IPC socket.
func GetIPCSocket() (ipc *IPCSocket, err error) {
	var out bytes.Buffer
	ipc = &IPCSocket{}

	cmd := exec.Command("i3", "--get-socketpath")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}

	path := strings.TrimSpace(out.String())
	sock, err := net.Dial("unix", path)
	ipc.socket = sock
	ipc.open = true
	return
}

// Receive a raw json bytestring from the socket and return a Message.
func (socket *IPCSocket) recv() (msg Message, err error) {
	header := make([]byte, _Headerlen)
	n, err := socket.socket.Read(header)

	// Check if this is a valid i3 message.
	if n != _Headerlen || err != nil {
		return
	}
	magicString := string(header[:len(_Magic)])
	if magicString != _Magic {
		err = MessageError(fmt.Sprintf(
			"Invalid magic string: got %q, expected %q.",
			magicString, _Magic))
		return
	}

	var bytelen [4]byte
	// Copy the byte values from the slice into the byte array. This is
	// necessary because the address of a slice does not point to the actual
	// values in memory.
	for i, b := range header[len(_Magic) : len(_Magic)+4] {
		bytelen[i] = b
	}
	length := *(*int32)(unsafe.Pointer(&bytelen))

	msg.Payload = make([]byte, length)
	n, err = socket.socket.Read(msg.Payload)
	if n != int(length) || err != nil {
		return
	}

	// Figure out the type of message.
	var bytetype [4]byte
	for i, b := range header[len(_Magic)+4 : len(_Magic)+8] {
		bytetype[i] = b
	}
	messageType := *(*uint32)(unsafe.Pointer(&bytetype))

	// Reminder: event messages have the highest bit of the type set to 1
	if messageType>>31 == 1 {
		msg.IsEvent = true
	}
	// Use the remaining bits
	msg.Type = int32(messageType & 0x7F)

	return
}

// Raw sends raw messages to i3. Returns a json byte string.
func (socket *IPCSocket) Raw(messageType MessageType, args string) (jsonReply []byte, err error) {
	// Set up the parts of the message.
	var (
		message  = []byte(_Magic)
		payload  = []byte(args)
		length   = int32(len(payload))
		bytelen  [4]byte
		bytetype [4]byte
	)

	// Black Magic™.
	bytelen = *(*[4]byte)(unsafe.Pointer(&length))
	bytetype = *(*[4]byte)(unsafe.Pointer(&messageType))

	for _, b := range bytelen {
		message = append(message, b)
	}
	for _, b := range bytetype {
		message = append(message, b)
	}
	for _, b := range payload {
		message = append(message, b)
	}

	_, err = socket.socket.Write(message)
	if err != nil {
		return
	}

	msg, err := socket.recv()
	if err == nil {
		jsonReply = msg.Payload
	}
	if msg.IsEvent {
		err = MessageTypeError("Received an event instead of a reply.")
	}
	return
}
