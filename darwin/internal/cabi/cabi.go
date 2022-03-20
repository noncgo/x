//go:build darwin && amd64
// +build darwin,amd64

// Package cabi provides utilities for calling C ABI functions without Cgo.
//
// It is a foreign function library in pure Go and allows calling functions in
// DLLs or shared libraries. It can be used to wrap these libraries in pure Go.
//
// Supported Platforms
//
// Currently the implementation supports the following GOOS/GOARCH combinations:
//  darwin/amd64
//
// Garbage Collector
//
// While currently Go has a non-moving GC, it’s possible that a copying GC will
// be implemented in the future. This package was written with that in mind and,
// after minor changes, should be compatible with copying GC, assuming that
// appropriate argument types were used for function calls in user’s code.
//
// Data Types
//
// This package defines a number of primitive C-compatible data types. The
// following table shows how Go types map to C types.
//  ┌───────────────┬────────────────┬────────────────────────┬─────────┐
//  │ Type          │ Go             │ C                      │ Usage   │
//  ├───────────────┼────────────────┼────────────────────────┼─────────┤
//  │ Void          │                │ void                   │     Out │
//  │ UnsafePointer │ unsafe.Pointer │ T* / uintptr_t         │ Arg     │
//  │ String        │ string         │ char*                  │ Arg     │
//  │ Bytes         │ []byte         │ char*                  │ Arg     │
//  │ Uintptr       │ uintptr        │ T* / uintptr_t         │ Arg     │
//  │ Bool          │ bool           │ _Bool                  │ Arg Out │
//  │ Int           │ int            │ long / ssize_t         │ Arg Out │
//  │ Int8          │ int8           │ char                   │ Arg Out │
//  │ Int16         │ int16          │ short                  │ Arg Out │
//  │ Int32         │ int32          │ int                    │ Arg Out │
//  │ Int64         │ int64          │ long long              │ Arg Out │
//  │ Uint          │ uint           │ unsigned long / size_t │ Arg Out │
//  │ Uint8         │ uint8          │ unsigned char          │ Arg Out │
//  │ Uint16        │ uint16         │ unsigned short         │ Arg Out │
//  │ Uint32        │ uint32         │ unsigned int           │ Arg Out │
//  │ Uint64        │ uint64         │ unsigned long long     │ Arg Out │
//  │ Float32       │ float32        │ float                  │ Arg     │
//  │ Float64       │ float64        │ double                 │ Arg     │
//  └───────────────┴────────────────┴────────────────────────┴─────────┘
// Note that, due to the variety of C compiler implementations, this may not
// apply to all platforms. In particular, long type is platform-specific, i.e.
// Windows uses LLP64 scheme while Darwin and Linux are LP64.
//
// Not all types can be used as a function outputs. For example, while it may be
// nice to get String and Bytes return values, both types would have to assume a
// certain ownership model and null-terminated memory. This assumption does not
// apply to all C functions and is likely confusing in many cases.
//
// Uintptr should used for pointers to unmanaged memory, while UnsafePointer
// must be used for pointers to Go values managed by GC. If the value escapes
// function call, caller must ensure that GC will not move the object or assume
// that GC is non-moving by importing go4.org/unsafe/assume-no-moving-gc package
// to avoid issues with future Go releases.
//
// String and Bytes arguments are passed as a pointer to the underlying data.
// Prefer using them to manually casting string or slice to runtime header.
//
// String passes a pointer to Go string representation. That is, null terminator
// is not appended and no copying is performed. It is invalid to pass String to
// functions that modify string data. If a function expects a null-terminated C
// string, use cstr.CString and pass the result as an UnsafePointer argument.
//
// Prior Art
//
// While the implementation is not a derivative of any other project, it shares
// the goal with libffi library and should have a similar interface. That’s not
// intentional though and there may be some subtle differences in the behavior.
//  https://github.com/libffi/libffi
//
package cabi

// Call invokes fn with the given arguments and expected output value.
func Call(fn uintptr, out Out, args ...Arg) {
	callg(fn, out, args)
}
