package dyld

import (
	"errors"
	"runtime"
)

const (
	// Each external function reference is bound the first time
	// the function is called.
	//
	// This is the default behavior.
	//
	BindLazy = _RTLD_LAZY

	// All external function references are bound immediately during
	// the call to Open.
	//
	// Lazy is normally preferred, for reasons of efficiency. However,
	// Now is useful to ensure that any undefined symbols are discovered.
	//
	BindNow = _RTLD_NOW

	// One of the following may be ORed into the mode argument:

	// Exported symbols will be available to any images built with
	// -flat_namespace option to ld(1) or to calls to Lookup functions.
	//
	// This is the default behavior.
	//
	ScopeGlobal = _RTLD_GLOBAL

	// Exported symbols are generally hidden and only availble to Lookup
	// when directly using the Dylib returned by Open.
	//
	ScopeLocal = _RTLD_LOCAL

	// The image is not loaded. However, a valid
	// Dylib is returned if the image already exists in the
	// process. This provides a way to query if an image is
	// already loaded. You eventually need a corresponding
	// call to Close.
	//
	NoLoad = _RTLD_NOLOAD

	// The image will never be removed from the address space, even after
	// all clients have released it via Close.
	//
	NoDelete = _RTLD_NODELETE

	// Additionally, the following may be ORed into the mode argument:

	// Lookup calls will only search the image specified, and not
	// subsequent images.
	//
	LookupFirst = _RTLD_FIRST
)

const (
	_RTLD_LAZY     = 0x1
	_RTLD_NOW      = 0x2
	_RTLD_LOCAL    = 0x4
	_RTLD_GLOBAL   = 0x8
	_RTLD_NOLOAD   = 0x10
	_RTLD_NODELETE = 0x80
	_RTLD_FIRST    = 0x100
)

var (
	MainOnly = _RTLD_MAIN_ONLY // Search main executable only.

	_RTLD_NEXT      = &Dylib{"", ^Handle(0)}
	_RTLD_DEFAULT   = &Dylib{"", ^Handle(1)}
	_RTLD_SELF      = &Dylib{"", ^Handle(2)}
	_RTLD_MAIN_ONLY = &Dylib{"", ^Handle(4)}
)

// LookupGlobal searches symbol by name in all Mach-O images in the process
// (except those loaded with ScopeLocal) in the order they were loaded. This
// can be a costly search and should be avoided.
//
// See dlsym(3).
// See RTLD_DEFAULT.
//
func LookupGlobal(name string) (*Symbol, error) {
	return _RTLD_DEFAULT.Lookup(name)
}

// LookupNext searches symbol by name in Mach-O images that were loaded after
// the one issuing this call.
//
// See dlsym(3).
// See RTLD_NEXT.
//
func LookupNext(name string) (*Symbol, error) {
	return _RTLD_NEXT.Lookup(name)
}

// LookupSelf is like LookupNext but also searches current Mach-O image.
//
// See dlsym(3).
// See RTLD_SELF.
//
func LookupSelf(name string) (*Symbol, error) {
	return _RTLD_SELF.Lookup(name)
}

// LookupMain searches the Mach-O image of the main executable.
//
// See dlsym(3).
// See RTLD_MAIN_ONLY.
//
func LookupMain(name string) (*Symbol, error) {
	return _RTLD_MAIN_ONLY.Lookup(name)
}

type Handle uintptr

type Dylib struct {
	Name   string
	Handle Handle
}

type Symbol struct {
	Dylib *Dylib
	Name  string
	addr  uintptr
}

// Addr returns the address of the symbol represented by s.
func (s *Symbol) Addr() uintptr { return s.addr }

func (s *Symbol) Invoke(args ...uintptr) (r1, r2, err uintptr) {
	// Hack for the example.
	if len(args) != 3 {
		panic("oops")
	}
	a1, a2, a3 := args[0], args[1], args[2]
	return syscall_syscall(s.addr, a1, a2, a3)
}

// MustOpen is like Open but panics if operation failes.
//
func MustOpen(path string, mode int) *Dylib {
	l, e := Open(path, mode)
	if e != nil {
		panic(e)
	}
	return l
}

// Open loads and links a dynamic library or bundle into the current process.
//
// Open examines the Mach-O file specified by path. If the file is compatible
// with the current process and has not already been loaded, it is loaded and
// linked. After being linked, if it contains any initializer functions, they
// are called, before Open returns.
//
// Open searches for a compatible Mach-O file in the directories specified by
// a set of environment variables and the process’s current working directory.
// When set, the environment variables must contain a colon-separated list of
// directory paths, which can be absolute or relative to the current working
// directory. The environment variables are LD_LIBRARY_PATH, DYLD_LIBRARY_PATH,
// and DYLD_FALLBACK_LIBRARY_PATH. The first two variables have no default
// value. The default value of DYLD_FALLBACK_LIBRARY_PATH is $HOME/lib;/usr/local/lib;/usr/lib.
// Open searches the directories specified in the environment variables in the
// order they are listed.
//
// When path doesn’t contain a slash character (i.e. it is just a leaf name),
// Open searches the following the locations until it finds a compatible Mach-O
// file: $LD_LIBRARY_PATH, $DYLD_LIBRARY_PATH, current working directory,
// $DYLD_FALLBACK_LIBRARY_PATH.
//
// When path contains a slash (i.e. a full path or a partial path), Open
// searches the following the locations until it finds a compatible Mach-O
// file: $DYLD_LIBRARY_PATH (with leaf name from path), current working
// directory (for partial paths), $DYLD_FALLBACK_LIBRARY_PATH (with leaf
// name from path).
//
// Note: There are no configuration files that control dlopen searching.
//
// Note: If the main executable is a set[ug]id binary, then all environment
// variables are ignored, and only a full path can be used.
//
// Note: macOS uses "universal" files to combine multiarch libraries. This also
// means that there are no separate 32-bit and 64-bit search paths.
//
// See dlopen(3).
//
func Open(path string, mode int) (*Dylib, error) {
	p, err := cstring(path)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	handle := dlopen(p, mode)
	if handle <= 0 {
		return nil, LastError()
	}
	h := Handle(handle)
	d := &Dylib{path, h}
	runtime.SetFinalizer(d, (*Dylib).Close)
	return d, nil
}

// IsLoadable preflights the load of a dynamic library or bundle.
//
// Returns true if the Mach-O file is compatible. If the file is not compatible,
// it returns false and an error.
//
// IsLoadable examines the Mach-O file specified by path. It checks if the file
// and libraries it depends on are all compatible with the current process.
// That is, they contain the correct architecture and are not otherwise ABI
// incompatible.
//
// IsLoadable uses the same steps as Open to find a compatible Mach-O file.
//
// See dlopen_preflight(3).
//
func IsLoadable(path string) (bool, error) {
	p, err := cstring(path)
	if err != nil {
		return false, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ok := dlopen_preflight(p)
	if !ok {
		err = LastError()
	}
	return ok, err
}

// Lookup searches symbol with name.
//
// Returns the address of the code or data location specified by the symbol name.
//
// See dlsym(3).
//
func (d *Dylib) Lookup(name string) (*Symbol, error) {
	p, err := cstring(name)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlsym(uintptr(d.Handle), p)
	// We must check dlerror because symbol could be NULL.
	if err := LastError(); err != nil {
		return nil, err
	}
	return &Symbol{d, name, ret}, nil
}

// Close closes a dynamic library or bundle.
//
// Close releases a reference to the dynamic library or bundle. If the reference
// count drops to 0, the bundle is removed from the address space. Just before
// removing a dynamic library or bundle in this way, any termination routines in
// it are called.
//
// There are a couple of cases in which a dynamic library will never be unloaded:
//
//   1) The main executable links against it,
//   2) An API that does not supoort unloading (e.g. NSAddImage()) was used
//      to load it or some other dynamic library that depends on it,
//   3) the dynamic library is in dyld’s shared cache.
//
// See dlclose(3).
//
func (d *Dylib) Close() (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlclose(uintptr(d.Handle))
	if ret != 0 {
		err = LastError()
	}

	// No need for a finalizer anymore.
	runtime.SetFinalizer(d, nil)
	return err
}

// LastError returns an error describing the last dyld error that occurred
// on this thread. At each call to LastError, the error indication is reset.
// Thus in the case of two calls to LastError, where the second call follows
// the first immediately, the second call will always return nil.
//
// See dlerror(3).
//
func LastError() error {
	ret := dlerror()
	if ret != 0 {
		s := gostring(ret)
		err := errors.New(s)
		return err
	}
	return nil
}
