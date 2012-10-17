package i3ipc

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	for _, s := range eventSockets {
		if !s.open {
			t.Error("Init failed: closed event socket found.")
		}
	}
	if len(eventSockets) != int(eventmax) {
		t.Errorf("Too much or not enough event sockets. Got %d, expected %d.\n",
			len(eventSockets), int(eventmax))
	}

	_, err := Subscribe(I3WorkspaceEvent)
	if err != nil {
		t.Errorf("Failed to subscribe: %f\n")
	}
	// TODO: A test to ensure that subscriptions work as intended.
}
