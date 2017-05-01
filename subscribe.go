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
	"log"
)

// EventType for subscribable events.
type EventType int32

// Enumeration of currently available event types.
const (
	I3WorkspaceEvent EventType = iota
	I3OutputEvent
	I3ModeEvent
	// private value used for setting up internal stuff in init()
	// The idea is that if there's a new type of event added to i3, it only
	// needs to be added here and in the payloads slice below, and the rest of
	// the code won't need to change.
	eventmax
)

// This slice is used to map event types to their string representation.
var payloads = []string{"workspace", "output", "mode"}

// AddEventType dynamically adds an event type by defining a name for it.
// Just in case i3 adds a new one and this library hasn't been updated yet.
// Returns the EventType which gets assigned to it.
//
// XXX: If you use this to add more than one new event type, add them in the
// RIGHT ORDER. I hope this case never pops up (because that would mean that
// this library is severely outdated), but I thought I'd warn you anyways.
func AddEventType(name string) (eventType EventType) {
	payloads = append(payloads, name)
	return EventType(len(payloads) - 1)
}

// Event describes an event reply from i3.
type Event struct {
	Type EventType
	// "change" is the name of the single field of the JSON map that i3 sends
	// when an event occurs, describing what happened.
	Change string
}

// Struct for replies from subscribe messages.
type subscribeReply struct {
	Success bool
}

// SubscribeError represents a subscription-related error.
type SubscribeError string

func (subscribeError SubscribeError) Error() string {
	return string(subscribeError)
}

// Private subscribe function. Sets up the socket.
func (socket *IPCSocket) subscribe(eventType EventType) (err error) {
	jsonReply, err := socket.Raw(I3Subscribe, "[\""+payloads[eventType]+"\"]")
	if err != nil {
		return
	}

	var subsReply subscribeReply
	err = json.Unmarshal(jsonReply, &subsReply)
	if err != nil {
		return
	}

	if !subsReply.Success {
		// TODO: Better error description.
		err = SubscribeError("Could not subscribe.")
	}
	return
}

// Subscribe to an event type. Returns a channel from which events can be read.
func Subscribe(eventType EventType) (subs chan Event, err error) {
	if eventType >= eventmax || eventType < 0 {
		err = SubscribeError("No such event type.")
		return
	}
	subs = make(chan Event)
	eventSockets[eventType].subscribers = append(
		eventSockets[eventType].subscribers, subs)
	return
}

// Listen for events on this socket, multiplexing them to all subscribers.
//
// XXX: This will cause all messages which are not events to be DROPPED.
func (socket *IPCSocket) listen() {
	for {
		if !socket.open {
			break
		}
		msg, err := socket.recv()
		// XXX: This ignores all errors. Maybe a FIXME, maybe not.
		if err != nil {
			continue
		}
		// Drop non-event messages.
		if !msg.IsEvent {
			continue
		}

		var event Event
		event.Type = EventType(msg.Type)
		json.Unmarshal(msg.Payload, &event)

		// Send each subscriber the event in a nonblocking manner.
		for _, subscriber := range socket.subscribers {
			select {
			case subscriber <- event: // NOP
			default:
				// If the event can't be written, just ignore this
				// subscriber.
			}
		}
	}
}

var eventSockets []*IPCSocket

// StartEventListener makes the library listen to events on the i3 socket
func StartEventListener() {
	// Check whether we have as much payloads as we have event types. You know,
	// just in case I'm coding on my third Club-Mate at 0400 in the morning when
	// updating this lib.
	if len(payloads) != int(eventmax) {
		log.Fatalf("Too much or not enough payloads: got %d, expected %d.\n",
			len(payloads), int(eventmax))
	}

	// Set up an IPCSocket to receive events for every type of event.
	var ev EventType
	for ; ev < eventmax; ev++ {
		sock, err := GetIPCSocket()
		if err != nil {
			log.Fatalf("Can't get i3 socket. Please make sure i3 is running. %v.", err)
		}
		err = sock.subscribe(ev)
		if err != nil {
			log.Fatalf("Can't subscribe: %v", err)
		}
		go sock.listen()
		if err != nil {
			log.Fatalf("Can't set up event sockets: %v", err)
		}

		eventSockets = append(eventSockets, sock)
	}
}
