package i3ipc

import (
	"testing"
)

func TestGetOutputs(t *testing.T) {
	ipc, _ := GetIPCSocket()

	//outputs, err := GetOutputs(ipc)
	_, err := ipc.GetOutputs()
	if err != nil {
		t.Errorf("Getting output list failed: %v", err)
	}
}
