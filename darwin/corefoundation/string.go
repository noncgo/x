//go:build darwin
// +build darwin

package corefoundation

// This file provides CFString APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfstring-rfh

import (
	"unsafe"

	_ "go4.org/unsafe/assume-no-moving-gc"

	"github.com/noncgo/x/darwin/internal/cabi"
	"github.com/noncgo/x/darwin/internal/cstr"
	"github.com/noncgo/x/darwin/internal/types"
)

// String is an opaque reference to a CFString type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfstringref
type String types.CFString

// StringEncoding specifies external string encodings.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfstringencoding
//  • https://developer.apple.com/documentation/corefoundation/cfstringbuiltinencodings
//  • https://developer.apple.com/documentation/corefoundation/cfstringencodings
type StringEncoding uint32

const (
	// StringEncodingUTF8 is an encoding constant that identifies the
	// UTF-8 encoding.
	//
	// References
	//  • https://developer.apple.com/documentation/corefoundation/cfstringbuiltinencodings/kcfstringencodingutf8
	StringEncodingUTF8 StringEncoding = 0x08000100
)

// UnsafeStr returns a pointer that can be used as a String value for the given
// Go string without calling into C code.
//
// Given s string len(s) <= math.MaxUint32 must be true. That is, string’s
// length must fit into 32 bits.
//
// Note that, unlike CFSTR macro, UnsafeStr may allocate memory. To avoid
// unnecessary allocations, store the result of UnsafeStr in a global variable.
//
// An example of valid usage for a hypothetical say function:
//
//  var str = corefoundation.UnsafeStr("hello")
//
//  func SayHello() {
//      cabi.Call(
//          abi.FuncPCABI0(say_trampoline),
//          cabi.Void(),
//          cabi.UnsafePointer(str),
//      )
//  }
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfstr
func UnsafeStr(s string) unsafe.Pointer {
	// assume-no-moving-gc: we should pin the given string.
	// TODO: panic on length overflow?
	ptr, length := cstr.Unpack(&s)
	type constString struct {
		isa    uintptr
		info   [4]byte
		rc     uint32
		ptr    unsafe.Pointer
		length uint32
	}
	return unsafe.Pointer(&constString{
		isa:    extern___CFConstantStringClassReference_getAddr(),
		info:   [4]byte{0xC8, 0x07},
		ptr:    ptr,
		length: uint32(length),
	})
}

// CreateArrayBySeparatingStrings an array of String object that represent
// substrings of the given s string separated by sep.
//
// If there was a problem creating the object, it returns false.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1541864-cfstringcreatearraybyseparatings
func CreateArrayBySeparatingStrings(alloc Allocator, s, sep String) (Array, bool) {
	var out uintptr
	cabi.Call(
		extern_CFStringCreateArrayBySeparatingStrings_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(alloc.Pointer()),
		cabi.Uintptr(s.Pointer()),
		cabi.Uintptr(sep.Pointer()),
	)
	return types.Pointer(out), out != 0
}

// CreateStringExternalRepresentation returns a Data object that stores the
// characters of the given String object as an “external representation”.
//
// It returns false if no loss byte was specified and the function could not
// convert the characters to the specified encoding, or if there was a problem
// creating the object.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1543169-cfstringcreateexternalrepresenta
func CreateStringExternalRepresentation(alloc Allocator, s String, enc StringEncoding, lossByte byte) (Data, bool) {
	var out uintptr
	cabi.Call(
		extern_CFStringCreateExternalRepresentation_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(alloc.Pointer()),
		cabi.Uintptr(s.Pointer()),
		cabi.Uint32(uint32(enc)),
		cabi.Uint8(lossByte),
	)
	return types.Pointer(out), out != 0
}

// CreateStringWithBytes creates a string from a buffer containing characters in
// a specified encoding.
//
// If the characters in the byte buffer are in an “external representation”
// format—that is, if the buffer contains a BOM (byte order marker), ext should
// be true. This is usually the case for bytes that are read in from a text file
// or received over the network. Otherwise, pass false.
//
// If there was a problem creating the object, it returns false.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/1543419-cfstringcreatewithbytes
func CreateStringWithBytes(alloc Allocator, data []byte, enc StringEncoding, ext bool) (String, bool) {
	var out uintptr
	cabi.Call(
		extern_CFStringCreateWithBytes_trampolineABI0,
		cabi.OutUintptr(&out),
		cabi.Uintptr(alloc.Pointer()),
		cabi.Bytes(data),
		cabi.Int(len(data)),
		cabi.Uint32(uint32(enc)),
		cabi.Bool(ext),
	)
	return types.Pointer(out), out != 0
}
