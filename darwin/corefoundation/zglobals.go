// Code generated by go run zgen.go. DO NOT EDIT.

//go:build darwin
// +build darwin

package corefoundation

import (
	"sync"

	"github.com/noncgo/x/darwin/internal/dyld"
)

// Uh, apparently cgo:cgo_import_dynamic links against function call stub?
// As a workaround, we use dlsym to get the right address from the loaded image.

//go:cgo_import_dynamic _ _ "/System/Library/Frameworks/CoreFoundation.framework/Versions/A/CoreFoundation"

var globals struct {
	kCFAllocatorDefault_addr uintptr
	kCFAllocatorDefault_once sync.Once

	kCFAllocatorMalloc_addr uintptr
	kCFAllocatorMalloc_once sync.Once

	kCFAllocatorMallocZone_addr uintptr
	kCFAllocatorMallocZone_once sync.Once

	kCFAllocatorNull_addr uintptr
	kCFAllocatorNull_once sync.Once

	kCFAllocatorSystemDefault_addr uintptr
	kCFAllocatorSystemDefault_once sync.Once

	kCFAllocatorUseContext_addr uintptr
	kCFAllocatorUseContext_once sync.Once

	kCFRunLoopCommonModes_addr uintptr
	kCFRunLoopCommonModes_once sync.Once

	kCFRunLoopDefaultMode_addr uintptr
	kCFRunLoopDefaultMode_once sync.Once

	kCFTypeArrayCallBacks_addr uintptr
	kCFTypeArrayCallBacks_once sync.Once

	__CFConstantStringClassReference_addr uintptr
	__CFConstantStringClassReference_once sync.Once
}

func extern_kCFAllocatorDefault_getAddr() uintptr {
	globals.kCFAllocatorDefault_once.Do(func() {
		sym, err := dyld.Lookup("kCFAllocatorDefault")
		if err != nil {
			panic(err)
		}
		globals.kCFAllocatorDefault_addr = sym.Addr
	})
	return globals.kCFAllocatorDefault_addr
}

func extern_kCFAllocatorMalloc_getAddr() uintptr {
	globals.kCFAllocatorMalloc_once.Do(func() {
		sym, err := dyld.Lookup("kCFAllocatorMalloc")
		if err != nil {
			panic(err)
		}
		globals.kCFAllocatorMalloc_addr = sym.Addr
	})
	return globals.kCFAllocatorMalloc_addr
}

func extern_kCFAllocatorMallocZone_getAddr() uintptr {
	globals.kCFAllocatorMallocZone_once.Do(func() {
		sym, err := dyld.Lookup("kCFAllocatorMallocZone")
		if err != nil {
			panic(err)
		}
		globals.kCFAllocatorMallocZone_addr = sym.Addr
	})
	return globals.kCFAllocatorMallocZone_addr
}

func extern_kCFAllocatorNull_getAddr() uintptr {
	globals.kCFAllocatorNull_once.Do(func() {
		sym, err := dyld.Lookup("kCFAllocatorNull")
		if err != nil {
			panic(err)
		}
		globals.kCFAllocatorNull_addr = sym.Addr
	})
	return globals.kCFAllocatorNull_addr
}

func extern_kCFAllocatorSystemDefault_getAddr() uintptr {
	globals.kCFAllocatorSystemDefault_once.Do(func() {
		sym, err := dyld.Lookup("kCFAllocatorSystemDefault")
		if err != nil {
			panic(err)
		}
		globals.kCFAllocatorSystemDefault_addr = sym.Addr
	})
	return globals.kCFAllocatorSystemDefault_addr
}

func extern_kCFAllocatorUseContext_getAddr() uintptr {
	globals.kCFAllocatorUseContext_once.Do(func() {
		sym, err := dyld.Lookup("kCFAllocatorUseContext")
		if err != nil {
			panic(err)
		}
		globals.kCFAllocatorUseContext_addr = sym.Addr
	})
	return globals.kCFAllocatorUseContext_addr
}

func extern_kCFRunLoopCommonModes_getAddr() uintptr {
	globals.kCFRunLoopCommonModes_once.Do(func() {
		sym, err := dyld.Lookup("kCFRunLoopCommonModes")
		if err != nil {
			panic(err)
		}
		globals.kCFRunLoopCommonModes_addr = sym.Addr
	})
	return globals.kCFRunLoopCommonModes_addr
}

func extern_kCFRunLoopDefaultMode_getAddr() uintptr {
	globals.kCFRunLoopDefaultMode_once.Do(func() {
		sym, err := dyld.Lookup("kCFRunLoopDefaultMode")
		if err != nil {
			panic(err)
		}
		globals.kCFRunLoopDefaultMode_addr = sym.Addr
	})
	return globals.kCFRunLoopDefaultMode_addr
}

func extern_kCFTypeArrayCallBacks_getAddr() uintptr {
	globals.kCFTypeArrayCallBacks_once.Do(func() {
		sym, err := dyld.Lookup("kCFTypeArrayCallBacks")
		if err != nil {
			panic(err)
		}
		globals.kCFTypeArrayCallBacks_addr = sym.Addr
	})
	return globals.kCFTypeArrayCallBacks_addr
}

func extern___CFConstantStringClassReference_getAddr() uintptr {
	globals.__CFConstantStringClassReference_once.Do(func() {
		sym, err := dyld.Lookup("__CFConstantStringClassReference")
		if err != nil {
			panic(err)
		}
		globals.__CFConstantStringClassReference_addr = sym.Addr
	})
	return globals.__CFConstantStringClassReference_addr
}