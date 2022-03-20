//go:build darwin
// +build darwin

package corefoundation_test

import (
	"runtime"
	"testing"

	_ "go4.org/unsafe/assume-no-moving-gc"

	"github.com/noncgo/x/darwin/corefoundation"
	"github.com/noncgo/x/darwin/internal/cstr"
	"github.com/noncgo/x/darwin/internal/types"
)

func TestCreateStringWithBytesWithoutAllocator(t *testing.T) {
	_, ok := corefoundation.CreateStringWithBytes(
		corefoundation.AllocatorNone(),
		nil,
		corefoundation.StringEncodingUTF8,
		false,
	)
	if ok {
		t.Fatal("creating a string with no-op allocator should be impossible")
	}
}

func TestUnsafeStrExternalRepresentation(t *testing.T) {
	const cs = "example"
	us := corefoundation.UnsafeStr(cs)
	defer runtime.KeepAlive(us)

	// assume-no-moving-gc: passing us as uintptr. Needs runtime.Pinner.
	s := corefoundation.String(types.Pointer(us))

	data, ok := corefoundation.CreateStringExternalRepresentation(
		corefoundation.AllocatorDefault(),
		s,
		corefoundation.StringEncodingUTF8,
		0,
	)
	if !ok {
		t.Fatal("failed to create an external representation of a constant string")
	}
	defer corefoundation.Release(data)

	gs := cstr.GoStringN(
		corefoundation.GetDataPointer(data),
		corefoundation.GetDataLength(data),
	)
	if cs != gs {
		t.Fatalf("external representation differs from constant string (expected %q, got %q)", cs, gs)
	}
}
