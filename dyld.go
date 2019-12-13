package dyld

import "runtime"

type Library struct {
	handle uintptr
}

type Symbol struct {
	library *Library
	symbol  uintptr
}

func (s *Symbol) Invoke(a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	// Hack for the example.
	return syscall_syscall(s.symbol, a1, a2, a3)
}

func Open(path string, mode int) (*Library, error) {
	p, err := cstring(path)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	handle := dlopen(p, mode)
	if handle <= 0 {
		return nil, dlerrorErr()
	}
	return &Library{handle}, nil
}

func Loadable(path string) (bool, error) {
	p, err := cstring(path)
	if err != nil {
		return false, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ok := dlopen_preflight(p)
	if !ok {
		err = dlerrorErr()
	}
	return ok, err
}

func (l *Library) Lookup(symbol string) (*Symbol, error) {
	p, err := cstring(symbol)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlsym(l.handle, p)
	// We must check dlerror because symbol could be NULL.
	if err := dlerrorErr(); err != nil {
		return nil, err
	}
	return &Symbol{l, ret}, nil
}

func (l *Library) Close() (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlclose(l.handle)
	if ret != 0 {
		err = dlerrorErr()
	}
	return err
}
