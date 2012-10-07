package i3ipc

import (
	"encoding/json"
	"net"
)

type I3Version struct {
	Major          int32
	Minor          int32
	Patch          int32
	Human_Readable string
}



func GetVersion(ipc net.Conn) (version I3Version, err error) {
	json_reply, err := Raw(I3GetVersion, "", ipc)
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &version)
	return
}
