package i3

import (
	"bytes"
	"net"
	"os/exec"
	"strings"
)

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
