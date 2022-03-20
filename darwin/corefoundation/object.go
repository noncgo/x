//go:build darwin
// +build darwin

package corefoundation

// This file provides CFType APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cftype

import (
	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/types"
)

// TypeID identifies particular Core Foundation opaque types.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cftypeid
type TypeID uint

// HashCode identifies an object in a hashing structure.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfhashcode
type HashCode uint

// Object is an opaque reference to a Core Foundation object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cftyperef
type Object types.CFType

// GetAllocator returns the allocator used to allocate a Core Foundation object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521280-cfgetallocator
func GetAllocator(v Object) Allocator {
	var out uintptr
	cabi.Call(
		extern_CFGetAllocator_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(v.Pointer()),
	)
	return types.Pointer(out)
}

// GetRetainCount returns the reference count of a Core Foundation object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521288-cfgetretaincount
func GetRetainCount(v Object) int {
	var out int
	cabi.Call(
		extern_CFGetRetainCount_trampolineABI0,
		cabi.OutInt(&out),
		cabi.Uintptr(v.Pointer()),
	)
	return out
}

// Equal determines whether two Core Foundation objects are considered equal.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521287-cfequal
func Equal(a, b Object) bool {
	var out bool
	cabi.Call(
		extern_CFEqual_trampolineABI0,
		cabi.OutBool(&out),
		cabi.Uintptr(a.Pointer()),
		cabi.Uintptr(b.Pointer()),
	)
	return out
}

// Hash returns a code that can be used to identify an object in a hashing
// structure.
//
// Two objects that are equal (as determined by the Equal function) have the
// same hashing value. However, the converse is not true: two objects with the
// same hashing value might not be equal. That is, hashing values are not
// necessarily unique.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521137-cfhash
func Hash(v Object) HashCode {
	var out uint
	cabi.Call(
		extern_CFHash_trampolineABI0,
		cabi.OutUint(&out),
		cabi.Uintptr(v.Pointer()),
	)
	return HashCode(out)
}

// CopyDescription returns a textual description of a Core Foundation object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521252-cfcopydescription
func CopyDescription(v Object) String {
	var out uintptr
	cabi.Call(
		extern_CFCopyDescription_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(v.Pointer()),
	)
	return types.Pointer(out)
}

// CopyTypeIDDescription returns a textual description of a Core Foundation
// type, as identified by its type ID, which can be used when debugging.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521220-cfcopytypeiddescription
func CopyTypeIDDescription(v TypeID) String {
	var out uintptr
	cabi.Call(
		extern_CFCopyTypeIDDescription_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uint(uint(v)),
	)
	return types.Pointer(out)
}

// GetTypeID returns the unique ID of an opaque type to which a Core Foundation
// object belongs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521218-cfgettypeid
func GetTypeID(v Object) TypeID {
	var out uint
	cabi.Call(
		extern_CFGetTypeID_trampolineABI0,
		cabi.OutUint(&out),
		cabi.Uintptr(v.Pointer()),
	)
	return TypeID(out)
}

// Show prints a description of a Core Foundation object to stderr.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1541433-cfshow
func Show(v Object) {
	cabi.Call(
		extern_CFShow_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(v.Pointer()),
	)
}

// Retain increments the reference counter of the given object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521269-cfretain
func Retain(v Object) Object {
	var out uintptr
	cabi.Call(
		extern_CFRetain_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(v.Pointer()),
	)
	return types.Pointer(out)
}

// Release decrements the reference counter of the given object. It releases
// the associated memory and resources if the counter reaches zero.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1521153-cfrelease/
func Release(v Object) {
	cabi.Call(
		extern_CFRelease_trampolineABI0,
		cabi.Void(),
		cabi.Uintptr(v.Pointer()),
	)
}
