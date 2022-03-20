//go:build darwin
// +build darwin

package fsevents

import (
	"github.com/noncgo/x/darwin/corefoundation"
)

// Event is the event passed to Callback.
//
// References
//  • https://developer.apple.com/documentation/coreservices/fseventstreamcallback
type Event struct {
	ID      EventID
	Path    string
	Flags   EventFlags
	ExtData corefoundation.Dictionary
}

// EventID is a filesystem change event ID.
//
// References
//  • https://developer.apple.com/documentation/coreservices/fseventstreameventid
type EventID uint64

const (
	// EventIDSinceNow is a special ID value that indicates to CreateStream
	// that the stream should resume from the latest event.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455359-anonymous/kfseventstreameventidsincenow
	//  • https://developer.apple.com/documentation/coreservices/1443980-fseventstreamcreate
	EventIDSinceNow EventID = 0xFFFFFFFFFFFFFFFF
)
