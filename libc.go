package dyld

import "unsafe"

func dlopen(path *byte, mode int) (ret uintptr) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlopen_trampoline), uintptr(unsafe.Pointer(path)), uintptr(mode), 0)
	ret = uintptr(r0)
	return
}

func dlopen_preflight(path *byte) (ret bool) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlopen_preflight_trampoline), uintptr(unsafe.Pointer(path)), 0, 0)
	ret = bool(r0 != 0)
	return
}

func dlerror() (ret uintptr) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlerror_trampoline), 0, 0, 0)
	ret = r0
	return
}

func dlclose(handle uintptr) (ret int) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlclose_trampoline), uintptr(handle), 0, 0)
	ret = int(r0)
	return
}

func dlsym(handle uintptr, symbol *byte) (ret uintptr) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlsym_trampoline), uintptr(handle), uintptr(unsafe.Pointer(symbol)), 0)
	ret = uintptr(r0)
	return
}

func libc_dlopen_trampoline()
func libc_dlopen_preflight_trampoline()
func libc_dlerror_trampoline()
func libc_dlclose_trampoline()
func libc_dlsym_trampoline()

//go:linkname libc_dlopen libc_dlopen
//go:linkname libc_dlopen_preflight libc_dlopen_preflight
//go:linkname libc_dlerror libc_dlerror
//go:linkname libc_dlclose libc_dlclose
//go:linkname libc_dlsym libc_dlsym

//go:cgo_import_dynamic libc_dlopen dlopen "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_dlopen_preflight dlopen_preflight "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_dlerror dlerror "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_dlclose dlclose "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_dlsym dlsym "/usr/lib/libSystem.B.dylib"
