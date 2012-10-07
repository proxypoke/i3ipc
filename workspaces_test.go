package i3ipc

import (
	"testing"
)

func TestGetWorkspaces(t *testing.T) {
	ipc, _ := GetIPCSocket()

	_, err := GetWorkspaces(ipc)
	if err != nil {
		t.Errorf("Getting workspace list failed: %v", err)
	}
}
