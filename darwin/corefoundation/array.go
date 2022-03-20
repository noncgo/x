//go:build darwin
// +build darwin

package corefoundation

// This file implements CFArray APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfarray-s28

import (
	"github.com/noncgo/x/darwin/internal/types"
)

// Array is an opaque reference to CFArray type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfarrayref
type Array types.CFArray

// ArrayCallbacks is a structure containing the callbacks of an Array.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfarraycallbacks
type ArrayCallbacks struct {
	// TODO: somehow wrap C functions that do accept context.
	// Perhaps we can use the Allocator’s context user info?
}

var arrayCallbacksForObject ArrayCallbacks

// ArrayCallbacksForObject returns an ArrayCallbacks appropriate for use when
// the values in an Array are all Core Foundation objects.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcftypearraycallbacks
func ArrayCallbacksForObject() *ArrayCallbacks {
	return &arrayCallbacksForObject
}
