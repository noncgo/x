// Package cstr provides utilities for interacting with C ABI functions that
// accept null-terminated strings.
package cstr

import (
	"strings"
	"unsafe"
)

// Unpack returns the underlying string data pointer and length.
func Unpack(s *string) (str unsafe.Pointer, len int) {
	// stringHeader is a safer version of reflect.StringHeader.
	type stringHeader struct {
		Data unsafe.Pointer
		Len  int
	}
	h := (*stringHeader)(unsafe.Pointer(s))
	return h.Data, h.Len
}

// CString returns a pointer to a copy of the string with null-terminator byte
// appended. It returns false if the given string already contains null byte.
func CString(s string) (*byte, bool) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, false
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return &a[0], true
}

// GoStringN copies n bytes of a C string  from unmanaged memory to GC-managed
// string. The returned string contains exactly n bytes.
func GoStringN(p uintptr, n int) (s string) {
	return runtime_gostringn(*(**byte)(unsafe.Pointer(&p)), n)
}

// GoString copies null-terminated C string from unmanaged memory to GC-managed
// string. The returned string does not contain null byte.
func GoString(p uintptr) string {
	return runtime_gostring(*(**byte)(unsafe.Pointer(&p)))
}

//go:linkname runtime_gostring runtime.gostring
func runtime_gostring(p *byte) string // from runtime/string.go

//go:linkname runtime_gostringn runtime.gostringn
func runtime_gostringn(p *byte, l int) string // from runtime/string.go
