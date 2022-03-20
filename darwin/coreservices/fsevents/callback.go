//go:build darwin
// +build darwin

package fsevents

import (
	"reflect"
	"unsafe"

	"github.com/noncgo/x/darwin/corefoundation"
	"github.com/noncgo/x/darwin/internal/cstr"
	"github.com/noncgo/x/darwin/internal/types"
)

// Callback is a callback function that will be called when FS events occur.
//
// This callback is invoked by the service from the client’s RunLoop(s) when
// events occur, per the parameters specified when the stream was created.
//
// References
//  • https://developer.apple.com/documentation/coreservices/fseventstreamcallback
type Callback func(stream ConstStream, info any, events ...Event)

type callbackInfo struct {
	// Callback is a Go function that callback calls back into from the C
	// callback (via crosscallCallback).
	Callback Callback

	// Info is the user info from stream’s context.
	Info any

	// Flags contains flags used to create the event stream.
	Flags CreateFlags
}

// callbackArgs carries arguments from the C callback to Callback function.
//
// It is allocated on the system stack in crosscallCallback.
//
//go:notinheap
type callbackArgs struct {
	stream     uintptr
	ctxt       *callbackInfo
	numEvents  uintptr
	eventPaths uintptr // CFArrayRef or **byte
	eventFlags uintptr // *EventFlags
	eventIDs   uintptr // *EventID
}

// callback invokes the Go actual callback passed via callbackInfo.
//
// Note that callback is invoked indirectly from runtime·cgocallbackg1, which
// uses ABIInternal calling convention.
//
// Currently it is not possible to select ABI implementation of a symbol in Go
// outside of the runtime package code, so we use a workaround similar to how we
// pass ABI0 wrappers via GLOBL/DATA, but this time from Go to assembly, to pass
// callback to CreateStream.
func callback(f *callbackArgs) {
	c := f.ctxt

	useCFTypes := 0 != c.Flags&CreateFlagUseCFTypes
	useExtData := 0 != c.Flags&CreateFlagUseExtendedData

	const (
		sizeofEventPath  = unsafe.Sizeof(uintptr(0))
		sizeofEventID    = unsafe.Sizeof(EventID(0))
		sizeofEventFlags = unsafe.Sizeof(EventFlags(0))
	)
	events := make([]Event, f.numEvents)
	for i := uintptr(0); i < f.numEvents; i++ {
		switch {
		case useCFTypes && useExtData:
			// TODO: f.eventPaths is CFArrayRef of CFDictionaryRef
			corefoundation.Show(types.Pointer(f.eventPaths))
		case useCFTypes:
			// TODO: f.eventPaths is CFArrayRef of CFStringRef
			corefoundation.Show(types.Pointer(f.eventPaths))
		default:
			s := cstr.GoString(f.eventPaths + i*unsafe.Sizeof(sizeofEventPath))
			events[i].Path = s
		}

		eventFlag := f.eventFlags + i*sizeofEventFlags
		events[i].Flags = **(**EventFlags)(unsafe.Pointer(&eventFlag))

		eventID := f.eventIDs + i*sizeofEventID
		events[i].ID = **(**EventID)(unsafe.Pointer(&eventID))
	}

	c.Callback(types.Pointer(f.stream), c.Info, events...)
}

// callbackABIInternal is a PC of an ABIInternal callback that we use in
// crosscallCallback.
//nolint:unused // used in assembly
var callbackABIInternal = reflect.ValueOf(callback).Pointer()

//nolint:unused // implemented in assembly
func crosscallCallback()

var crosscallCallbackABI0 uintptr
