package i3ipc

import (
	"testing"
)

func TestGetBarConfig(t *testing.T) {
	ipc, _ := GetIPCSocket()

	ids, err := GetBarIds(ipc)
	if err != nil {
		t.Errorf("Getting bar IDs failed: %v", err)
	}

	id := ids[0]
	//bar, err := GetBarConfig(ipc)
	_, err = GetBarConfig(ipc, id)
	if err != nil {
		t.Errorf("Getting bar config failed: %v", err)
	}
}
