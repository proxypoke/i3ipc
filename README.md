[![GoDoc](https://godoc.org/github.com/mdirkse/i3ipc?status.svg)](http://godoc.org/github.com/mdirkse/i3ipc/)
[![Build Status](https://travis-ci.org/mdirkse/i3ipc.svg?branch=master)](https://travis-ci.org/mdirkse/i3ipc)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdirkse/i3ipc)](https://goreportcard.com/report/github.com/mdirkse/i3ipc)

i3ipc
=====

Overview
--------
i3ipc is a golang library for convenient access to the IPC API of the [i3 window
manager](http://i3wm.org).

Capabilities
------------
As of the time of writing, this library is able to access all functionality of
the IPC API of i3. This includes sending commands and other message types, as
well as handling subscriptions.

If you just want a quick overview of the documentation, head to
[http://godoc.org/github.com/mdirkse/i3ipc/][doc].

Usage
-----
Thanks to Go's built-in git support, you can start using i3ipc with a simple

    import "github.com/mdirkse/i3ipc"

For everything except subscriptions, you will want to create an IPCSocket over
which the communication will take place. This object has methods for all message
types that i3 will accept, though some might be split into multiple methods (eg.
*Get_Bar_Config*). You can create such a socket quite easily:

    ipcsocket, err := i3ipc.GetIPCSocket()

As a simple example of what you could do next, let's get the version of i3 over
our new socket:

    version, err := ipcsocket.GetVersion()

For further commands, refer to `go doc` or use the aforementioned
[website][doc].

### Subscriptions
i3ipc handles subscriptions in a convenient way: you don't have to think about
managing the socket or watch out for unordered replies. The appropriate method
simply returns a channel from which you can read Event objects.

Here's a simple example - we start the event listener, we subscribe to workspace
events, then simple print all of them as we receive them:

    i3ipc.StartEventListener()
    ws_events, err := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
    for {
        event := <-ws_events
        fmt.Printf("Received an event: %v\n", event)
    }

i3ipc currently has no way of subscribing to multiple event types over a single
channel. If you want this, you can simply create multiple subscriptions, then
demultiplex those channels yourself - `select` is your friend.

[doc]: http://godoc.org/github.com/mdirkse/i3ipc/
