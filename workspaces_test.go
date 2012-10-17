package i3ipc

import (
	"testing"
)

func TestGetWorkspaces(t *testing.T) {
	ipc, _ := GetIPCSocket()

	_, err := ipc.GetWorkspaces()
	if err != nil {
		t.Errorf("Getting workspace list failed: %v", err)
	}
}
