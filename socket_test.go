package i3ipc

import (
	"testing"
)

func TestGetIPCSocket(t *testing.T) {
	ipc, err := GetIPCSocket()
	if err != nil {
		t.Errorf("Failed to acquire the IPC socket: %v", err)
	}
	ipc.Close()
	if ipc.open {
		t.Error("IPC socket appears open after closing.")
	}
}

func TestRaw(t *testing.T) {
	ipc, _ := GetIPCSocket()

	_, err := ipc.Raw(I3GetWorkspaces, "")
	if err != nil {
		t.Errorf("Raw message sending failed: %v", err)
	}
}
