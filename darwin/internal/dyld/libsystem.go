//go:build darwin
// +build darwin

package dyld

import (
	"unsafe"

	"github.com/noncgo/x/darwin/internal/cabi"
)

type dyldInfo struct {
	fname uintptr // const char*
	fbase uintptr // void*
	sname uintptr // const char*
	saddr uintptr // void*
}

func dlopen(path *byte, mode int) (ret uintptr) {
	cabi.Call(
		extern_dlopen_trampolineABI0,
		cabi.OutUintptr(&ret),
		cabi.UnsafePointer(unsafe.Pointer(path)),
		cabi.Int(mode),
	)
	return
}

func dlopen_preflight(path *byte) (ret bool) {
	cabi.Call(
		extern_dlopen_preflight_trampolineABI0,
		cabi.OutBool(&ret),
		cabi.UnsafePointer(unsafe.Pointer(path)),
	)
	return
}

func dlerror() (ret uintptr) {
	cabi.Call(
		extern_dlerror_trampolineABI0,
		cabi.OutUintptr(&ret),
	)
	return
}

func dlclose(handle uintptr) (ret int) {
	cabi.Call(
		extern_dlclose_trampolineABI0,
		cabi.OutInt(&ret),
		cabi.Uintptr(handle),
	)
	return
}

func dlsym(handle uintptr, symbol *byte) (ret uintptr) {
	cabi.Call(
		extern_dlsym_trampolineABI0,
		cabi.OutUintptr(&ret),
		cabi.Uintptr(handle),
		cabi.UnsafePointer(unsafe.Pointer(symbol)),
	)
	return
}

func dladdr(addr uintptr, info *dyldInfo) (ret int) {
	cabi.Call(
		extern_dladdr_trampolineABI0,
		cabi.OutInt(&ret),
		cabi.Uintptr(addr),
		cabi.UnsafePointer(unsafe.Pointer(info)),
	)
	return
}
