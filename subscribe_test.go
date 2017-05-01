// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i3ipc

import (
	"testing"
)

func TestInit(t *testing.T) {
	StartEventListener()

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
		t.Errorf("Failed to subscribe: %f\n", err)
	}
	// TODO: A test to ensure that subscriptions work as intended.
}
