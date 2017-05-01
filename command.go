// Author: slowpoke <mail plus git at slowpoke dot io>
// Repository: https://github.com/proxypoke/i3ipc
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License.
// It comes without any warranty, to the extent permitted by
// applicable law. For a copy of the license, see COPYING or
// head to http://sam.zoy.org/wtfpl/COPYING.

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
