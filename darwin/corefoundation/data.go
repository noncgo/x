//go:build darwin
// +build darwin

package corefoundation

// This file provides CFData APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfdata-rv9

import (
	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/types"
)

// Data is an opaque reference to a CFData type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfdata
type Data types.CFData

// GetDataPointer returns a read-only pointer to the bytes of a Data object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1543330-cfdatagetbyteptr
func GetDataPointer(d Data) uintptr {
	var out uintptr
	cabi.Call(
		extern_CFDataGetBytePtr_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(d.Pointer()),
	)
	return out
}

// GetDataLength returns the number of bytes contained by a Data object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1541728-cfdatagetlength
func GetDataLength(d Data) int {
	var out int
	cabi.Call(
		extern_CFDataGetLength_trampolineABI0,
		cabi.OutInt(&out),
		cabi.Uintptr(d.Pointer()),
	)
	return out
}
