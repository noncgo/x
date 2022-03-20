//go:build darwin
// +build darwin

// Package dyld provides a bindings for dynamic linker APIs.
package dyld

import (
	"errors"
	"runtime"
	"syscall"

	"github.com/noncgo/x/darwin/internal/cstr"
)

// ErrNotFound is an error returned from Addr function when there is no image
// for the given address.
var ErrNotFound = errors.New("no image found for the given address")

const (
	// BindLazy specifies the binding mode for loading an image.
	//
	// Each external function reference is bound the first time
	// the function is called.
	//
	// This is the default behavior.
	//
	BindLazy = rtldLazy

	// BindLazy specifies the binding mode for loading an image.
	//
	// All external function references are bound immediately
	// during the call to Open.
	//
	// Lazy is normally preferred, for reasons of efficiency.
	// However, BindNow is useful to ensure that any undefined
	// symbols are discovered.
	//
	BindNow = rtldNow

	// One of the following may be ORed into the mode argument:

	// ScopeGlobal specifies the scope of a loaded image.
	//
	// Exported symbols will be available to any images built with
	// -flat_namespace option to ld(1) or to calls to Lookup functions.
	//
	// This is the default behavior.
	//
	ScopeGlobal = rtldGlobal

	// ScopeLocal specifies the scope of a loaded image.
	//
	// Exported symbols are generally hidden and only available
	// to Lookup when directly using the Image returned by Open.
	//
	ScopeLocal = rtldLocal

	// NoLoad specifies the mode for loading an image.
	//
	// The image is not loaded. However, a valid Image is returned if
	// the image already exists in the process. This provides a way to
	// query if an image is already loaded.
	//
	// You eventually need a corresponding call to Close.
	//
	NoLoad = rtldNoLoad

	// NoDelete specifies the mode for loading an image.
	//
	// The image will never be removed from the address space,
	// even after all clients have released it via Close.
	//
	NoDelete = rtldNoDelete

	// Additionally, the following may be ORed into the mode argument:

	// LookupFirst specifies the mode for lookups in a loaded image.
	//
	// Lookup calls will only search the image specified, and not
	// subsequent images.
	//
	LookupFirst = rtldFirst
)

const (
	rtldLazy     = 0x1
	rtldNow      = 0x2
	rtldLocal    = 0x4
	rtldGlobal   = 0x8
	rtldNoLoad   = 0x10
	rtldNoDelete = 0x80
	rtldFirst    = 0x100
)

var (
	rtldNext     = &Image{"", ^Handle(0)}
	rtldDefault  = &Image{"", ^Handle(1)}
	rtldSelf     = &Image{"", ^Handle(2)}
	rtldMainOnly = &Image{"", ^Handle(4)}
)

// Lookup searches symbol by name in all Mach-O images in the process
// (except those loaded with ScopeLocal) in the order they were loaded. This
// can be a costly search and should be avoided.
//
// See dlsym(3).
// See RTLD_DEFAULT.
//
func Lookup(name string) (*Symbol, error) {
	return rtldDefault.Lookup(name)
}

// LookupNext searches symbol by name in Mach-O images that were loaded after
// the one issuing this call.
//
// See dlsym(3).
// See RTLD_NEXT.
//
func LookupNext(name string) (*Symbol, error) {
	return rtldNext.Lookup(name)
}

// LookupSelf is like LookupNext but also searches current Mach-O image.
//
// See dlsym(3).
// See RTLD_SELF.
//
func LookupSelf(name string) (*Symbol, error) {
	return rtldSelf.Lookup(name)
}

// LookupMain searches the Mach-O image of the main executable.
//
// See dlsym(3).
// See RTLD_MAIN_ONLY.
//
func LookupMain(name string) (*Symbol, error) {
	return rtldMainOnly.Lookup(name)
}

// Handle is an opaque handle for the loaded image of a dynamic library.
type Handle uintptr

// Image represents an image of a loaded dynamic library.
type Image struct {
	Name   string
	Handle Handle
}

// Symbol represents a symbol in a loaded dynamic library.
type Symbol struct {
	Image *Image
	Name  string
	Addr  uintptr
}

// MustOpen is like Open but panics if operation fails.
//
func MustOpen(path string, mode int) *Image {
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
// and DYLD_FALLBACK_LIBRARY_PATH. The default value of the latter variable is
// $HOME/lib;/usr/local/lib;/usr/lib. The first two variables have no default
// value. Open searches the directories specified in the environment variables
// in the order they are listed.
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
func Open(path string, mode int) (*Image, error) {
	p, err := cstring(path)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	handle := dlopen(p, mode)
	if handle <= 0 {
		return nil, lastError()
	}
	h := Handle(handle)
	d := &Image{path, h}
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
		err = lastError()
	}
	return ok, err
}

// Lookup searches symbol with name.
//
// Returns the address of the code or data location specified by the symbol name.
//
// See dlsym(3).
//
func (d *Image) Lookup(name string) (*Symbol, error) {
	p, err := cstring(name)
	if err != nil {
		return nil, err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlsym(uintptr(d.Handle), p)
	// We must check dlerror because symbol could be NULL.
	if err := lastError(); err != nil {
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
//   3) The dynamic library is in dyld’s shared cache.
//
// See dlclose(3).
//
func (d *Image) Close() (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := dlclose(uintptr(d.Handle))
	if ret != 0 {
		err = lastError()
	}

	// No need for a finalizer anymore.
	runtime.SetFinalizer(d, nil)
	return err
}

// AddrInfo describes an address in process address space.
type AddrInfo struct {
	Fname string  // Pathname of shared object
	Fbase uintptr // Base address of shared object
	Sname string  // Name of nearest symbol
	Saddr uintptr // Address of nearest symbol
}

// Addr finds the image containing a given address.
//
// Returns information about the image containing the address addr.
// If the image containing addr is found, but no nearest symbol was
// found, info.Sname and info.Saddr fields are set to zero value.
//
// See dladdr(3).
//
func Addr(p uintptr) (info *AddrInfo, err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var dli dyldInfo
	ret := dladdr(p, &dli)
	if ret != 0 {
		fname := gostring(dli.fname)
		sname := gostring(dli.sname)
		info := AddrInfo{
			Fname: fname,
			Fbase: dli.fbase,
			Sname: sname,
			Saddr: dli.saddr,
		}
		return &info, nil
	}
	return nil, ErrNotFound
}

// lastError returns an error describing the last dyld error that occurred
// on this thread. At each call to lastError, the error indication is reset.
// Thus in the case of two calls to lastError, where the second call follows
// the first immediately, the second call will always return nil.
//
// See dlerror(3).
//
func lastError() error {
	ret := dlerror()
	if ret != 0 {
		s := gostring(ret)
		// TODO export error vars by known string suffixes.
		// E.g. strings.HasSuffix(s, ": symbol not found").
		err := errors.New(s)
		return err
	}
	return nil
}

func cstring(s string) (*byte, error) {
	p, ok := cstr.CString(s)
	if !ok {
		return nil, syscall.EINVAL
	}
	return p, nil
}

func gostring(p uintptr) string {
	return cstr.GoString(p)
}
