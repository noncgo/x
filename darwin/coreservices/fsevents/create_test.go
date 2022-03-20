//go:build darwin
// +build darwin

package fsevents_test

import (
	"testing"
	"time"

	"github.com/noncgo/x/darwin/corefoundation"
	"github.com/noncgo/x/darwin/coreservices/fsevents"
)

func TestCreateStreamWithoutAllocator(t *testing.T) {
	noPaths, ok := corefoundation.CreateMutableArray(corefoundation.AllocatorDefault(), 0, nil)
	if !ok {
		t.Fatal("failed to create an empty array for testing")
	}
	defer corefoundation.Release(noPaths)

	_, ok = fsevents.CreateStream(
		corefoundation.AllocatorNone(),
		nil, // callback
		nil, // context
		noPaths,
		fsevents.EventIDSinceNow,
		time.Second,
		fsevents.CreateFlagNone,
	)
	if ok {
		t.Fatal("creating a stream with no-op allocator should be impossible")
	}
}
