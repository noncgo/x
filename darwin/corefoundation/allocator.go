//go:build darwin
// +build darwin

package corefoundation

// This file provides CFAllocator APIs.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfallocator

import (
	"unsafe"

	"github.com/noncgo/x/darwin/internal/types"
)

// Allocator is an opaque reference to a CFAllocator type.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/cfallocatorref
type Allocator types.CFAllocator

// AllocatorDefault returns an allocator that is synonym for NULL.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfallocatordefault
func AllocatorDefault() Allocator {
	addr := extern_kCFAllocatorDefault_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// AllocatorSystem returns default system allocator.
//
// Refererences
//  • https://developer.apple.com/documentation/corefoundation/kcfallocatorsystemdefault
func AllocatorSystem() Allocator {
	addr := extern_kCFAllocatorSystemDefault_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// AllocatorMalloc returns an allocator that uses malloc, realloc and free.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfallocatormalloc
func AllocatorMalloc() Allocator {
	addr := extern_kCFAllocatorMalloc_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// AllocatorMallocZone returns an allocator that explicitly uses the default
// malloc zone, returned by malloc_default_zone.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfallocatormalloczone
func AllocatorMallocZone() Allocator {
	addr := extern_kCFAllocatorMallocZone_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// AllocatorNone returns an allocator does not nothing.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfallocatornull
func AllocatorNone() Allocator {
	addr := extern_kCFAllocatorNull_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}

// AllocatorUseContext returns a special allocator argument to CreateAllocator
// that uses the functions given in the context to allocate the allocator.
//
// References
//  • https://developer.apple.com/documentation/corefoundation/kcfallocatorusecontext
func AllocatorUseContext() Allocator {
	addr := extern_kCFAllocatorUseContext_getAddr()
	addr = **(**uintptr)(unsafe.Pointer(&addr))
	return types.Pointer(addr)
}
