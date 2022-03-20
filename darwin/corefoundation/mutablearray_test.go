//go:build darwin
// +build darwin

package corefoundation_test

import (
	"testing"

	"github.com/noncgo/x/darwin/corefoundation"
)

func TestMutableArrayWithoutAllocator(t *testing.T) {
	_, ok := corefoundation.CreateMutableArray(
		corefoundation.AllocatorNone(),
		0,
		nil,
	)
	if ok {
		t.Fatal("creating an array with no-op allocator should be impossible")
	}
}

func TestMutableArrayWithCallbacksForObject(t *testing.T) {
	const capacity = 2
	a, ok := corefoundation.CreateMutableArray(
		corefoundation.AllocatorDefault(),
		capacity,
		corefoundation.ArrayCallbacksForObject(),
	)
	if !ok {
		t.Fatal("failed to create a mutable array")
	}
	defer func() {
		if a != nil {
			corefoundation.Release(a)
		}
	}()

	v, ok := corefoundation.CreateStringWithBytes(
		corefoundation.AllocatorDefault(),
		[]byte("hello world"),
		corefoundation.StringEncodingUTF8,
		false,
	)
	if !ok {
		t.Fatal("failed to create a string")
	}
	defer corefoundation.Release(v)

	corefoundation.AppendToArray(a, v)
	if corefoundation.GetRetainCount(v) != 2 {
		t.Fatal("appending to array with retain callback must increment element refcount")
	}

	corefoundation.Release(a)
	if corefoundation.GetRetainCount(v) != 1 {
		t.Fatal("releasing array with release callback must decrement element refcount")
	}

	// Do not release array in the deferred function.
	a = nil
}
