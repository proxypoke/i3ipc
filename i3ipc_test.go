package i3ipc

import (
	"testing"
)

func TestGetIPCSocket(t *testing.T) {
	ipc, err := GetIPCSocket()
	defer ipc.Close()
	if err != nil {
		t.Errorf("Failed to acquire the IPC socket: %v", err)
	}
}

func TestRaw(t *testing.T) {
	ipc, _ := GetIPCSocket()

	_, err := Raw(I3GetWorkspaces, "", ipc)
	if err != nil {
		t.Errorf("Raw message sending failed: %v", err)
	}
}
