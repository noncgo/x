//go:build darwin
// +build darwin

package dyld_test

import (
	"fmt"

	"github.com/noncgo/x/darwin/internal/dyld"
)

func ExampleLookup_scopeGlobal() {
	sym, err := dyld.Lookup("dlopen")
	if err != nil {
		fmt.Printf("lookup: %v", err)
		return
	}

	fmt.Println("Found", sym.Name)
	// Output:
	// Found dlopen
}

func ExampleAddr_lookup() {
	sym, err := dyld.Lookup("dlopen")
	if err != nil {
		fmt.Printf("lookup: %v", err)
		return
	}

	info, err := dyld.Addr(sym.Addr)
	if err != nil {
		fmt.Printf("addr: %v", err)
		return
	}

	fmt.Println("Found", info.Sname, "in", info.Fname)
	// Output:
	// Found dlopen in /usr/lib/system/libdyld.dylib
}
