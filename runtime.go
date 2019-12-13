package dyld

import _ "unsafe"

//go:linkname syscall_syscall syscall.syscall
//go:linkname runtime_gostring runtime.gostring
//go:linkname funcPC runtime.funcPC

func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) // runtime/sys_darwin.go
func runtime_gostring(p *byte) string                              // runtime/string.go
func funcPC(f interface{}) uintptr                                 // runtime/proc.go
