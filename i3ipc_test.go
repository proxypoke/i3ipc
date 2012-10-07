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

func TestCommand(t *testing.T) {
	ipc, _ := GetIPCSocket()
	defer ipc.Close()

	// `exec /bin/true` is a good NOP operation for testing
	success, err := Command("exec /bin/true", ipc)
	if !success {
		t.Error("Unsuccessful command.")
	}
	if err != nil {
		t.Errorf("An error occurred during command: %v", err)
	}
}
