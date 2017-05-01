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

// I3Bar represents the configuration of a bar. For documentation of the
// fields, refer to http://i3wm.org/docs/ipc.html#_bar_config_reply.
type I3Bar struct {
	ID               string
	Mode             string
	Position         string
	StatusCommand    string
	Font             string
	WorkspaceButtons bool
	Verbose          bool
	Colors           Colors
}

// Colors represents colors as used in I3Bar.
type Colors struct {
	FocusedWorkspaceBorder string
	FocusedWorkspaceBg     string
	FocusedWorkspaceText   string
}

// GetBarIds fetches a list of IDs for all configured bars.
func (socket *IPCSocket) GetBarIds() (ids []string, err error) {
	jsonReply, err := socket.Raw(I3GetBarConfig, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &ids)
	return
}

// GetBarConfig returns the configuration of the bar with the given ID.
func (socket *IPCSocket) GetBarConfig(id string) (bar I3Bar, err error) {
	jsonReply, err := socket.Raw(I3GetBarConfig, id)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &bar)
	return
}
