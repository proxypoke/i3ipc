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
	"encoding/json"
)

// CommandError for replies from a command to i3.
type CommandError string

func (error CommandError) Error() string {
	return string(error)
}

// Struct for replies from command messages.
type commandReply struct {
	Success bool
	Error   string
}

// Command sends a command to i3.
// FIXME: Doesn't support chained commands yet.
func (socket *IPCSocket) Command(action string) (success bool, err error) {
	jsonReply, err := socket.Raw(I3Command, action)
	if err != nil {
		return
	}

	var cmdReply []commandReply
	err = json.Unmarshal(jsonReply, &cmdReply)
	if err != nil {
		return
	}

	success = cmdReply[0].Success
	if cmdReply[0].Error == "" {
		err = nil
	} else {
		err = CommandError(cmdReply[0].Error)
	}

	return
}
