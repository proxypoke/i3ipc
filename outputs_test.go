package i3ipc

import (
	"testing"
)

func TestGetOutputs(t *testing.T) {
	ipc, _ := GetIPCSocket()

	//outputs, err := GetOutputs(ipc)
	_, err := GetOutputs(ipc)
	if err != nil {
		t.Errorf("Getting output list failed: %v", err)
	}
}
