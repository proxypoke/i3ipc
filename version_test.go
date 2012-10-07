package i3ipc

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	ipc, _ := GetIPCSocket()

	_, err := GetVersion(ipc)
	if err != nil {
		t.Errorf("Getting version failed: %v", err)
	}
}
