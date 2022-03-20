//go:build darwin
// +build darwin

// The fsevents command monitors changes in the given directory trees using
// macOS File System Events framework.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/noncgo/x/darwin/corefoundation"
	"github.com/noncgo/x/darwin/coreservices/fsevents"
)

var (
	latency     = flag.Duration("latency", time.Second, "")
	since       = flag.Uint64("since", uint64(fsevents.EventIDSinceNow), "")
	useCFTypes  = flag.Bool("use-cftypes", true, "")
	noDefer     = flag.Bool("no-defer", false, "")
	watchRoot   = flag.Bool("watch-root", false, "")
	ignoreSelf  = flag.Bool("ignore-self", false, "")
	fileEvents  = flag.Bool("file-events", true, "")
	markSelf    = flag.Bool("mark-self", false, "")
	useExtData  = flag.Bool("use-ext-data", true, "")
	fullHistory = flag.Bool("full-history", false, "")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	paths, ok := corefoundation.CreateMutableArray(
		corefoundation.AllocatorDefault(),
		len(args),
		corefoundation.ArrayCallbacksForObject(),
	)
	if !ok {
		log.Fatal("failed to create paths array")
	}
	for _, arg := range args {
		v, ok := corefoundation.CreateStringWithBytes(
			corefoundation.AllocatorDefault(),
			[]byte(arg),
			corefoundation.StringEncodingUTF8,
			false,
		)
		if !ok {
			log.Fatal("failed to create path string")
		}
		corefoundation.AppendToArray(paths, v)
		corefoundation.Release(v)
	}
	stream, ok := fsevents.CreateStream(
		corefoundation.AllocatorDefault(),
		func(s fsevents.ConstStream, info any, events ...fsevents.Event) {
			for _, ev := range events {
				fmt.Printf("%+v\n", ev)
			}
		},
		nil,
		paths,
		fsevents.EventID(*since),
		*latency,
		streamFlags(),
	)
	if !ok {
		log.Fatal("failed to create filesystem event stream")
	}
	corefoundation.Release(paths)

	// Spawn a new thread to get a run loop. Most native applications will
	// have run loop on a main thread.
	//
	// This example demonstrates how to set up a disposable run loop.
	runLoop, stopRunLoop := createRunLoop()

	fsevents.ScheduleStreamWithRunLoop(
		stream,
		runLoop,
		corefoundation.RunLoopDefaultMode(),
	)
	corefoundation.WakeUpRunLoop(runLoop)
	if !fsevents.StartStream(stream) {
		log.Fatal("failed to start filesystem event stream")
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	<-sigc

	fsevents.StopStream(stream)
	fsevents.InvalidateStream(stream)
	fsevents.ReleaseStream(stream)

	corefoundation.Release(runLoop)

	stopRunLoop()
}

// createRunLoop creates a new background thread and returns its run loop.
func createRunLoop() (corefoundation.RunLoop, func()) {
	c := make(chan corefoundation.RunLoop, 1)
	go func() {
		// Taint thread to release all associated resources when
		// goroutine exits. In particular, this should release
		// current thread’s CFRunLoopRef instance.
		//
		// See also https://github.com/golang/go/issues/20458
		runtime.LockOSThread()

		runLoop := corefoundation.GetCurrentRunLoop()
		_ = corefoundation.Retain(runLoop)
		c <- runLoop
		corefoundation.RunCurrentRunLoop()
		close(c)
	}()
	runLoop := <-c
	return runLoop, func() {
		corefoundation.WakeUpRunLoop(runLoop)
		corefoundation.StopRunLoop(runLoop)
		<-c // wait until goroutine exits
	}
}

func streamFlags() fsevents.CreateFlags {
	var v fsevents.CreateFlags
	if *useCFTypes {
		v |= fsevents.CreateFlagUseCFTypes
	}
	if *noDefer {
		v |= fsevents.CreateFlagNoDefer
	}
	if *watchRoot {
		v |= fsevents.CreateFlagWatchRoot
	}
	if *ignoreSelf {
		v |= fsevents.CreateFlagIgnoreSelf
	}
	if *fileEvents {
		v |= fsevents.CreateFlagFileEvents
	}
	if *markSelf {
		v |= fsevents.CreateFlagMarkSelf
	}
	if *useExtData {
		v |= fsevents.CreateFlagUseExtendedData
	}
	if *fullHistory {
		v |= fsevents.CreateFlagFullHistory
	}
	return v
}
