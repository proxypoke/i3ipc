package i3

import (
	"testing"
)


func TestGetIPCSocket(t *testing.T) {
	ipc, err := GetIPCSocket()
	if err != nil {
		t.Errorf("%v", err)
	}
	_, err = ipc.Write([]byte("i3-ipc"))
	if err != nil {
		t.Errorf("Failed to write to the IPC socket: %v", err)
	}

	reply := make([]byte, 1024)
	_, err =ipc.Read(reply)
	if err != nil {
		t.Errorf("Failed to read from the IPC socket: %v", err)
	}

	ipc.Close()
}
