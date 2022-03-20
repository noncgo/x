//go:build darwin
// +build darwin

package fsevents

import (
	"time"
	"unsafe"

	_ "go4.org/unsafe/assume-no-moving-gc"

	"github.com/noncgo/x/darwin/corefoundation"
	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/types"
)

// privateStream is the Stream interface implementation that keeps callback
// user info alive.
type privateStream struct {
	types.FSEventStreamRef
	CallbackInfo *callbackInfo
}

// newStream returns a new Stream instance.
func newStream(v uintptr, c *callbackInfo) *privateStream {
	return &privateStream{
		FSEventStreamRef: types.Pointer(v),
		CallbackInfo:     c,
	}
}

// CreateStream creates a new FS event stream object with the given parameters.
//
// References
//  • https://developer.apple.com/documentation/coreservices/1443980-fseventstreamcreate
func CreateStream(
	alloc corefoundation.Allocator,
	callback Callback,
	context *StreamContext,
	paths corefoundation.Array,
	sinceWhen EventID,
	latency time.Duration,
	flags CreateFlags,
) (Stream, bool) {
	type streamContext struct {
		Version         int // zero
		Info            *callbackInfo
		Retain          uintptr
		Release         uintptr
		CopyDescription uintptr
	}

	// assume-no-moving-gc: info pointer to escapes to unmanaged memory.
	//
	// TODO: put in global sync.Map and set Retain/Release callbacks in
	// context.
	//
	// Currently we use privateStream to keep info alive, but that breaks
	// comparison for Stream and ConstStream values.
	//
	info := &callbackInfo{
		Callback: callback,
		Flags:    flags,
	}
	if context != nil {
		info.Info = context.Info
	}
	ctxt := &streamContext{Info: info}

	var out uintptr
	cabi.Call(
		extern_FSEventStreamCreate_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(alloc.Pointer()),
		cabi.Uintptr(crosscallCallbackABI0),
		cabi.UnsafePointer(unsafe.Pointer(ctxt)),
		cabi.Uintptr(paths.Pointer()),
		cabi.Uint64(uint64(sinceWhen)),
		cabi.Float64(latency.Seconds()),
		cabi.Uint32(uint32(flags)),
	)
	return newStream(out, info), out != 0
}
