package main

import (
	"log"

	"github.com/tie/dyld"
)

func main() {
	lib, err := dyld.Open("/System/Library/Frameworks/AppKit.framework/AppKit", 0)
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer func() {
		err := lib.Close()
		if err != nil {
			log.Printf("close: %v", err)
		}
	}()

	sym, err := lib.Lookup("NSApplicationMain")
	if err != nil {
		log.Fatalf("lookup: %v", err)
	}

	sym.Invoke(0, 0, 0)

	log.Println(lib, sym)

}
