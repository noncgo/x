//go:build darwin
// +build darwin

package fsevents

import (
	"fmt"
	"strings"
)

// CreateFlags is a set of flags that modify the behavior of the FS event stream
// being created.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags
//  • https://developer.apple.com/documentation/coreservices/fseventstreamcreateflags
type CreateFlags uint32

const (
	// CreateFlagNone is the default mode for CreateStream.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagnone
	CreateFlagNone CreateFlags = 0

	// CreateFlagUseCFTypes indicates that the framework should invoke
	// callback with CF types rather than raw C types.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagusecftypes
	CreateFlagUseCFTypes CreateFlags = 1 << (iota - 1)

	// CreateFlagNoDefer affects the meaning of the latency parameter. If
	// this flag is set and more than latency seconds have elapsed since the
	// last event, an application will receive the event immediately. The
	// delivery of the event resets the latency timer and any further events
	// will be delivered after latency seconds have elapsed.
	//
	// Unless this flag is set, then when an event occurs after a period of
	// no events, the latency timer is started. Any events that occur during
	// the next latency seconds will be delivered as one group (including
	// that first event). The delivery of the group of events resets the
	// latency timer and any further events will be delivered after latency
	// seconds.
	//
	// This flag is useful for apps that are interactive and want to react
	// immediately to changes but avoid getting swamped by notifications
	// when changes are occurring in rapid succession. The default behavior
	// is more appropriate for background, daemon or batch processing apps.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagnodefer
	CreateFlagNoDefer

	// CreateFlagWatchRoot requests notifications of changes along the path
	// to the path(s) being monitored.
	//
	// For example, with this flag, if path “/foo/bar” is monitored and it
	// is renamed to “/foo/bar.old”, an EventFlagRootChanged event is sent
	// to the application. The same is true if the directory “/foo” was
	// renamed. The event sent in this case is a special event: the path for
	// the event is the original path specified in during stream creation,
	// the flag EventFlagRootChanged is set and event ID is zero.
	//
	// These events are useful to indicate that a particular hierarchy
	// should be rescanned because it changed completely (as opposed to the
	// things inside of it changing).
	//
	// To track the location of a changed directory, it is best to open the
	// directory before creating the stream and find the current path via a
	// file descriptor.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagwatchroot
	CreateFlagWatchRoot

	// CreateFlagIgnoreSelf indicates that events triggered by the current
	// process should not be sent.
	//
	// Note that this has no effect on historical events, i.e., those
	// delivered before the EventFlagHistoryDone sentinel event.
	//
	// Also, this does not apply to EventFlagRootChanged events because the
	// CreateFlagWatchRoot feature uses a separate mechanism that is unable
	// to provide information about the responsible process.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagignoreself
	CreateFlagIgnoreSelf

	// CreateFlagFileEvents requests file-level notifications.
	//
	// With this flag set, stream will receive events about individual files
	// in the hierarchy being monitored instead of only receiving directory
	// level notifications.
	//
	// Use this flag with care as it will generate significantly more events
	// than without it.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagfileevents
	CreateFlagFileEvents

	// CreateFlagMarkSelf indicates that events triggered by the current
	// process should have EventFlagOwnEvent flag.
	//
	// Note that this has no effect on historical events, i.e., those
	// delivered before the EventFlagHistoryDone sentinel event.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagmarkself
	CreateFlagMarkSelf

	// CreateFlagUseExtendedData indicates that the framework should pass
	// CFArrayRef of CFDictionaryRefs to the callback function instead of
	// CFArrayRef of CFStringRefs.
	//
	// Requires CreateFlagUseCFTypes flag to be set.
	//
	// See the EventExtendedData*Key definitions for the set of keys that
	// may be set in the dictionary.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflaguseextendeddata
	CreateFlagUseExtendedData

	// CreateFlagFullHistory indicates that all historical events in a given
	// chunk should be returned even if their event ID is less than the
	// sinceWhen ID. Otherwise, when requesting historical events, it is
	// possible that some events may get skipped due to the way they are
	// stored.
	//
	// In other words, it delivers all the events in the first chunk of
	// historical events that contains the sinceWhen ID so that none are
	// skipped even if their ID is less than the sinceWhen ID. This overlap
	// avoids any issue with missing events that happened at/near the time
	// of an unclean restart of the client process.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/1455376-fseventstreamcreateflags/kfseventstreamcreateflagfullhistory
	CreateFlagFullHistory
)

// String implements the fmt.Stringer interface.
func (k CreateFlags) String() string {
	if name, ok := k.bitString(); ok {
		return name
	}
	var names []string
	for i := CreateFlags(1); i != 0 && i <= k; i <<= 1 {
		if k&i == 0 {
			continue
		}
		name, ok := i.bitString()
		if !ok {
			continue
		}
		k ^= i
		names = append(names, name)
	}
	if k != 0 {
		names = append(names, fmt.Sprintf("%b", k))
	}
	return strings.Join(names, "|")
}

func (k CreateFlags) bitString() (string, bool) {
	var v string
	switch k {
	case CreateFlagNone:
		v = "None"
	case CreateFlagUseCFTypes:
		v = "UseCFTypes"
	case CreateFlagNoDefer:
		v = "NoDefer"
	case CreateFlagWatchRoot:
		v = "WatchRoot"
	case CreateFlagIgnoreSelf:
		v = "IgnoreSelf"
	case CreateFlagFileEvents:
		v = "FileEvents"
	case CreateFlagMarkSelf:
		v = "MarkSelf"
	case CreateFlagFullHistory:
		v = "FullHistory"
	case CreateFlagUseExtendedData:
		v = "UseExtendedData"
	default:
		return "", false
	}
	return v, true
}
