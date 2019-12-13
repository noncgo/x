package dyld

import (
	"strings"
	"syscall"
	"unsafe"
)

//go:linkname syscall_syscall syscall.syscall
//go:linkname runtime_gostring runtime.gostring
//go:linkname funcPC runtime.funcPC

func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) // runtime/sys_darwin.go
func runtime_gostring(p *byte) string                              // runtime/string.go
func funcPC(f interface{}) uintptr                                 // runtime/proc.go

func cstring(s string) (*byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, syscall.EINVAL
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return &a[0], nil
}

func gostring(p uintptr) string {
	return runtime_gostring(*(**byte)(unsafe.Pointer(&p)))
}
