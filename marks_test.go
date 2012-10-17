package i3ipc

import (
	"testing"
)

func TestGetMarks(t *testing.T) {
	ipc, _ := GetIPCSocket()

	_, err := ipc.GetMarks()
	if err != nil {
		t.Errorf("Getting marks failed: %v", err)
	}
}
