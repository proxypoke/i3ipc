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
	"encoding/json"
)

// I3Version represents the version of i3. For documentation of the
// fields, refer to http://i3wm.org/docs/ipc.html#_version_reply.
type I3Version struct {
	Major                int32
	Minor                int32
	Patch                int32
	HumanReadable        string
	LoadedConfigFileName string
}

// GetVersion fetches the version of i3.
func (socket *IPCSocket) GetVersion() (version I3Version, err error) {
	jsonReply, err := socket.Raw(I3GetVersion, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &version)
	return
}
