//go:build darwin
// +build darwin

package corefoundation

// This file provides CFMutableArray APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfmutablearray-rrk

import (
	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/types"
)

// MutableArray is an opaque reference to CFMutableArray type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfmutablearrayref
type MutableArray types.CFMutableArray

// CreateMutableArray creates a new mutable array.
//
// If there was a problem creating the object, it returns false.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1388770-cfarraycreatemutable
func CreateMutableArray(alloc Allocator, capacity int, callbacks *ArrayCallbacks) (MutableArray, bool) {
	var callbacksAddr uintptr
	if callbacks == ArrayCallbacksForObject() {
		callbacksAddr = extern_kCFTypeArrayCallBacks_getAddr()
	}

	var out uintptr
	cabi.Call(
		extern_CFArrayCreateMutable_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(alloc.Pointer()),
		cabi.Int(capacity),
		cabi.Uintptr(callbacksAddr),
	)
	return types.Pointer(out), out != 0
}

// AppendToArray adds a value to an array giving it the new largest index.
//
// It is invalid to append an element to an array with non-zero capacity beyond
// the set limit or zero limit if the caller cannot guarantee that the system
// has enough memory to accommodate a new element.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1388802-cfarrayappendvalue
func AppendToArray(m MutableArray, v Object) {
	var objectAddr uintptr
	if v != nil {
		objectAddr = v.Pointer()
	}

	cabi.Call(
		extern_CFArrayAppendValue_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(m.Pointer()),
		cabi.Uintptr(objectAddr),
	)
}
