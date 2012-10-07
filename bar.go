package i3ipc

import (
	"encoding/json"
	"net"
)

type I3Bar struct {
	Id                string
	Mode              string
	Position          string
	Status_command    string
	Font              string
	Workspace_Buttons bool
	Verbose           bool
	Colors            Colors
}

type Colors struct {
	Focused_Workspace_Border string
	Focused_Workspace_Bg     string
	Focused_Workspace_Text   string
}

func GetBarIds(ipc net.Conn) (ids []string, err error) {
	json_reply, err := Raw(ipc, I3GetBarConfig, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &ids)
	return
}

func GetBarConfig(ipc net.Conn, id string) (bar I3Bar, err error) {
	json_reply, err := Raw(ipc, I3GetBarConfig, id)
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &bar)
	return
}
