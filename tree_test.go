package i3ipc

import (
	"testing"
)

func TestGetTree(t *testing.T) {
	ipc, _ := GetIPCSocket()

	//root, err := GetTree(ipc)
	_, err := ipc.GetTree()
	if err != nil {
		t.Errorf("Getting tree failed: %v", err)
	}
}
