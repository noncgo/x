//go:build darwin
// +build darwin

package fsevents

import (
	"fmt"
	"strings"
)

// EventFlags represents a set of flags that can be passed to Callback function.
//
// References
//  • https://developer.apple.com/documentation/coreservices/file_system_events/1455361-fseventstreameventflags
//  • https://developer.apple.com/documentation/coreservices/1455361-fseventstreameventflags
type EventFlags uint32

// TODO: copy-paste docs for remaining constants

const (
	// EventFlagNone indicates that there was some change in the directory
	// at the specific path supplied in this event.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflagnone
	EventFlagNone EventFlags = 0

	// EventFlagMustScanSubDirs indicates that application must rescan not
	// just the directory given in the event, but all its children,
	// recursively.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflagmustscansubdirs
	EventFlagMustScanSubDirs EventFlags = 1 << (iota - 1)

	// EventFlagUserDropped may be set in addition to
	// EventFlagMustScanSubDirs to indicate that a problem occurred in
	// buffering the events.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflaguserdropped
	EventFlagUserDropped

	// EventFlagKernelDropped may be set in addition to
	// EventFlagMustScanSubDirs to indicate that a problem occurred in
	// buffering the events.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflagkerneldropped
	EventFlagKernelDropped

	// EventFlagEventIDsWrapped indicates that 64-bit event ID counter
	// wrapped around.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflageventidswrapped
	EventFlagEventIDsWrapped

	// EventFlagHistoryDone denotes a sentinel event sent to mark the end of
	// the “historical” events sent as a result of specifying a sinceWhen in
	// the CreateStream call. It will not be sent if EventIDSinceNow was
	// passed for sinceWhen.
	//
	// The client should ignore the path supplied in this callback.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflaghistorydone
	EventFlagHistoryDone

	// EventFlagRootChanged denotes a special event sent when there is a
	// change to one of the directories along the path to one of the
	// directories being monitored.
	//
	// When this flag is set, the event ID is zero and the path corresponds
	// to one of the paths being monitored (specifically, the one that
	// changed). The path may no longer exist because it or one of its
	// parents was deleted or renamed.
	//
	// Events with this flag set will only be sent if the
	// CreateFlagWatchRoot flag was passed to CreateStream.
	//
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflagrootchanged
	EventFlagRootChanged

	// EventFlagMount denotes a special event sent when a volume is mounted
	// underneath one of the paths being monitored.
	//
	// The path in the event is the path to the newly-mounted volume.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflagmount
	EventFlagMount

	// EventFlagUnmount denotes a special event sent when a volume is
	// unmounted underneath one of the paths being monitored.
	//
	// The path in the event is the path to the directory from which the
	// volume was unmounted.
	//
	// References
	//  • https://developer.apple.com/documentation/coreservices/kfseventstreameventflagunmount
	EventFlagUnmount

	EventFlagItemCreated
	EventFlagItemRemoved
	EventFlagItemInodeMetaMod
	EventFlagItemRenamed
	EventFlagItemModified
	EventFlagItemFinderInfoMod
	EventFlagItemChangeOwner
	EventFlagItemXattrMod
	EventFlagItemIsFile
	EventFlagItemIsDir
	EventFlagItemIsSymlink
	EventFlagOwnEvent
	EventFlagItemIsHardlink
	EventFlagItemIsLastHardlink
	EventFlagItemCloned
)

// String implements the fmt.Stringer interface.
func (k EventFlags) String() string {
	if name, ok := k.bitString(); ok {
		return name
	}
	var names []string
	for i := EventFlags(1); i != 0 && i <= k; i <<= 1 {
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

func (k EventFlags) bitString() (string, bool) {
	var v string
	switch k {
	case EventFlagNone:
		v = "None"
	case EventFlagMustScanSubDirs:
		v = "MustScanSubDirs"
	case EventFlagUserDropped:
		v = "UserDropped"
	case EventFlagKernelDropped:
		v = "KernelDropped"
	case EventFlagEventIDsWrapped:
		v = "EventIDsWrapped"
	case EventFlagHistoryDone:
		v = "HistoryDone"
	case EventFlagRootChanged:
		v = "RootChanged"
	case EventFlagMount:
		v = "Mount"
	case EventFlagUnmount:
		v = "Unmount"
	case EventFlagItemCreated:
		v = "ItemCreated"
	case EventFlagItemRemoved:
		v = "ItemRemoved"
	case EventFlagItemInodeMetaMod:
		v = "ItemInodeMetaMod"
	case EventFlagItemRenamed:
		v = "ItemRenamed"
	case EventFlagItemModified:
		v = "ItemModified"
	case EventFlagItemFinderInfoMod:
		v = "ItemFinderInfoMod"
	case EventFlagItemChangeOwner:
		v = "ItemChangeOwner"
	case EventFlagItemXattrMod:
		v = "ItemXattrMod"
	case EventFlagItemIsFile:
		v = "ItemIsFile"
	case EventFlagItemIsDir:
		v = "ItemIsDir"
	case EventFlagItemIsSymlink:
		v = "ItemIsSymlink"
	case EventFlagOwnEvent:
		v = "OwnEvent"
	case EventFlagItemIsHardlink:
		v = "ItemIsHardlink"
	case EventFlagItemIsLastHardlink:
		v = "ItemIsLastHardlink"
	case EventFlagItemCloned:
		v = "ItemCloned"
	default:
		return "", false
	}
	return v, true
}
