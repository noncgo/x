package dyld_test

import (
	"fmt"

	"github.com/tie/dyld"
)

func ExampleOpen_scopeLocal() {
	lib, err := dyld.Open("libsqlite3.dylib", dyld.ScopeLocal)
	if err != nil {
		fmt.Printf("open: %v", err)
		return
	}

	sym, err := lib.Lookup("sqlite3_open")
	if err != nil {
		fmt.Printf("lookup: %v", err)
		return
	}

	// The symbols from dylib are not globally visible.
	if sym, err := dyld.Lookup(sym.Name); err == nil {
		fmt.Printf("Oops, symbol %s is global!", sym.Name)
		return
	}

	fmt.Println("Found", sym.Name)
	// Output:
	// Found sqlite3_open
}

func ExampleLookup() {
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
