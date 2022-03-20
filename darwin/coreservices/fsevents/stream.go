//go:build darwin
// +build darwin

package fsevents

import (
	"github.com/noncgo/x/darwin/corefoundation"
	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/types"
)

// Stream is an opaque reference to a FSEventStream type.
//
// References
//  • https://developer.apple.com/documentation/coreservices/fseventstreamref
type Stream types.FSEventStreamRef

// ConstStream is an opaque reference to a constant FSEventStream type.
//
// References
//  • https://developer.apple.com/documentation/coreservices/constfseventstreamref
type ConstStream types.ConstFSEventStreamRef

// StreamContext is a structure containing stream’s user info.
//
// References
//  • https://developer.apple.com/documentation/coreservices/fseventstreamcontext
type StreamContext struct {
	// Info is an arbitrary client-defined value to be associated with the
	// stream and passed to the callback when it is invoked.
	Info any

	// TODO: also add allocator callbacks.
}

// ShowStream prints a description of the supplied stream to stderr for
// debugging purposes.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1444302-fseventstreamshow
func ShowStream(s ConstStream) {
	cabi.Call(
		extern_FSEventStreamShow_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
	)
}

// RetainStream increments the stream’s reference counter.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1444986-fseventstreamretain?language=objc
func RetainStream(s Stream) {
	cabi.Call(
		extern_FSEventStreamRetain_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
	)
}

// ReleaseStream decrements the stream’s reference counter. The counter is
// initially one and is incremented via RetainStream. If the counter reaches
// zero then the stream is deallocated.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1445989-fseventstreamrelease
func ReleaseStream(s Stream) {
	cabi.Call(
		extern_FSEventStreamRelease_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
	)
}

// StartStream attempts to register with the FS Events service to receive events
// per the parameters in the stream.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1448000-fseventstreamstart
func StartStream(s Stream) bool {
	var out bool
	cabi.Call(
		extern_FSEventStreamStart_trampolineABI0,
		cabi.OutBool(&out),
		cabi.Uintptr(s.Pointer()),
	)
	return out
}

// StopStream unregisters the stream from the FS Events service.
//
// Once stopped, the stream can be restarted via StartStream, at which point it
// will resume receiving events from where it left off (“sinceWhen”).
//
// References
//  • https://developer.apple.com/documentation/coreservices/1447673-fseventstreamstop
func StopStream(s Stream) {
	cabi.Call(
		extern_FSEventStreamStop_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
	)
}

// ScheduleStreamWithRunLoop schedules the stream on the specified run loop.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1447824-fseventstreamschedulewithrunloop
func ScheduleStreamWithRunLoop(s Stream, runLoop corefoundation.RunLoop, mode corefoundation.RunLoopMode) {
	cabi.Call(
		extern_FSEventStreamScheduleWithRunLoop_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
		cabi.Uintptr(runLoop.Pointer()),
		cabi.Uintptr(mode.Pointer()),
	)
}

// UnscheduleStreamFromRunLoop unschedules the stream from the specified run
// loop.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1441982-fseventstreamunschedulefromrunlo
func UnscheduleStreamFromRunLoop(s Stream, runLoop corefoundation.RunLoop, mode corefoundation.RunLoopMode) {
	cabi.Call(
		extern_FSEventStreamUnscheduleFromRunLoop_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
		cabi.Uintptr(runLoop.Pointer()),
		cabi.Uintptr(mode.Pointer()),
	)
}

// InvalidateStream unschedules the stream from any run loops or dispatch queues
// upon which it had been scheduled.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1446990-fseventstreaminvalidate
func InvalidateStream(s Stream) {
	cabi.Call(
		extern_FSEventStreamInvalidate_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(s.Pointer()),
	)
}
