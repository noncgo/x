//go:build darwin
// +build darwin

package corefoundation

// This file provides CFRunLoop APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfrunloop-rht

import (
	"time"
	"unsafe"

	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/types"
)

// RunLoop is an opaque reference to a CFRunLoop type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfrunloop-rht
//  • https://developer.apple.com/documentation/corefoundation/cfrunloopref
type RunLoop types.CFRunLoop

// RunLoopMode determines what events are processed by the run loop during a
// given iteration.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfrunloopmode
//  • https://developer.apple.com/documentation/corefoundation/cfrunloop-rht
type RunLoopMode = String

// RunLoopDefaultMode returns the default mode of the run loop.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfrunloopdefaultmode
//  • https://developer.apple.com/documentation/corefoundation/cfrunloopmode
//  • https://developer.apple.com/documentation/corefoundation/cfrunloop-rht
func RunLoopDefaultMode() RunLoopMode {
	addr := extern_kCFRunLoopDefaultMode_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// RunLoopCommonModes returns a special pseudo-mode that allows
// associating more than one mode with a given run loop source.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfrunloopcommonmodes
//  • https://developer.apple.com/documentation/corefoundation/cfrunloopmode
//  • https://developer.apple.com/documentation/corefoundation/cfrunloop-rht
func RunLoopCommonModes() RunLoopMode {
	addr := extern_kCFRunLoopCommonModes_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// RunLoopResult identifies the reason run loop exited RunCurrentRunLoopInMode.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfrunloop/cfrunloopruninmode_exit_codes
//  • https://developer.apple.com/documentation/corefoundation/cfrunlooprunresult
type RunLoopResult int32

const (
	_ RunLoopResult = iota

	// RunLoopResultFinished indicates that the running run loop mode has no
	// sources or timers to process.
	//
	// References
	//  • https://developer.apple.com/documentation/corefoundation/cfrunlooprunresult/finished
	RunLoopResultFinished

	// RunLoopResultStopped indicates that StopRunLoop was called on the run
	// loop.
	//
	// References
	//  • https://developer.apple.com/documentation/corefoundation/cfrunlooprunresult/stopped
	RunLoopResultStopped

	// RunLoopResultTimedOut indicates that the specified time interval for
	// running the run loop has passed.
	//
	// References
	//  • https://developer.apple.com/documentation/corefoundation/cfrunlooprunresult/timedout
	RunLoopResultTimedOut

	// RunLoopResultHandledSource indicates that a source has been processed
	// if the run loop was told to run only until a source was processed.
	//
	// References
	//  • https://developer.apple.com/documentation/corefoundation/cfrunlooprunresult/handledsource
	RunLoopResultHandledSource
)

// GetMainRunLoop returns the RunLoop object for the main thread.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1542890-cfrunloopgetmain
func GetMainRunLoop() RunLoop {
	var out uintptr
	cabi.Call(
		extern_CFRunLoopGetMain_trampolineABI0,
		cabi.OutUintptr(&out),
	)
	return types.Pointer(out)
}

// GetCurrentRunLoop returns the RunLoop object for the current thread.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1542428-cfrunloopgetcurrent
func GetCurrentRunLoop() RunLoop {
	var out uintptr
	cabi.Call(
		extern_CFRunLoopGetCurrent_trampolineABI0,
		cabi.OutUintptr(&out),
	)
	return types.Pointer(out)
}

// RunCurrentRunLoopInMode runs the current thread’s RunLoop object in a
// particular mode.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1541988-cfrunloopruninmode
func RunCurrentRunLoopInMode(mode RunLoopMode, d time.Duration, returnAfterSourceHandled bool) RunLoopResult {
	var out int32
	cabi.Call(
		extern_CFRunLoopRunInMode_trampolineABI0,
		cabi.OutInt32(&out),
		cabi.Uintptr(mode.Pointer()),
		cabi.Float64(d.Seconds()),
		cabi.Bool(returnAfterSourceHandled),
	)
	return RunLoopResult(out)
}

// RunCurrentRunLoop runs the current thread’s RunLoop object in its default
// mode indefinitely or until the run loop is stopped with StopRunLoop.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1542011-cfrunlooprun
func RunCurrentRunLoop() {
	cabi.Call(
		extern_CFRunLoopRun_trampolineABI0,
		cabi.Void(),
	)
}

// WakeUpRunLoop wakes a waiting RunLoop object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1541622-cfrunloopwakeup
func WakeUpRunLoop(r RunLoop) {
	cabi.Call(
		extern_CFRunLoopRun_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(r.Pointer()),
	)
}

// StopRunLoop forces a RunLoop object to stop running.
//
// Note that if the run loop is nested with a callout from one activation
// starting another activation running, only the innermost activation is exited.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1541796-cfrunloopstop
func StopRunLoop(r RunLoop) {
	cabi.Call(
		extern_CFRunLoopStop_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(r.Pointer()),
	)
}
