package dyld

import (
	"errors"
	"unsafe"
)

func dlerrorErr() error {
	ret := dlerror()
	if ret != 0 {
		s := uintptrToString(ret)
		return errors.New(s)
	}
	return nil
}

func uintptrToString(p uintptr) string {
	b := *(**byte)(unsafe.Pointer(&p))
	s := runtime_gostring(b)
	return s
}
