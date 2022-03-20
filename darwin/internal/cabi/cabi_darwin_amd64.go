package cabi

import (
	"math"
	"reflect"
	"runtime"
	"unsafe"

	_ "go4.org/unsafe/assume-no-moving-gc"
)

// TODO: support stack arguments

type frame struct {
	FuncPC uintptr

	DI, SI, DX, CX, R8, R9         uintptr
	AX                             uintptr
	X0, X1, X2, X3, X4, X5, X6, X7 uintptr
}

func callg(fn uintptr, out Out, args []Arg) {
	var f frame
	f.FuncPC = fn

	const (
		numGP = 6
		numFP = 8
	)
	var gp, fp uintptr
	for _, arg := range args {
		switch arg.typ {
		case argTypeVoid:
			continue
		case argTypeUnsafePointer,
			argTypeString,
			argTypeBytes,
			argTypeUintptr,
			argTypeBool,
			argTypeInt,
			argTypeInt8,
			argTypeInt16,
			argTypeInt32,
			argTypeInt64,
			argTypeUint,
			argTypeUint8,
			argTypeUint16,
			argTypeUint32,
			argTypeUint64:
			if gp == numGP {
				// TODO: push to the stack
				// continue
				panic("not implemented")
			}
			var v uintptr
			switch arg.typ {
			case argTypeUnsafePointer:
				// assume-no-moving-gc: either Go GC is non-moving
				// or pointer address should be pinned.
				//
				// Once runtime.Pinner gets implemented, we should be
				// safe against a potential moving GC in the future.
				//
				// See also https://github.com/golang/go/issues/46787
				v = uintptr(arg.getUnsafePointer())
			case argTypeString:
				// assume-no-moving-gc: same as above.
				p := arg.getString()
				v = (*reflect.StringHeader)(unsafe.Pointer(&p)).Data
			case argTypeBytes:
				// assume-no-moving-gc: same as above.
				p := arg.getBytes()
				v = (*reflect.SliceHeader)(unsafe.Pointer(&p)).Data
			case argTypeUintptr:
				v = arg.getUintptr()
			case argTypeBool:
				if arg.getBool() {
					v = 1
				}
			case argTypeInt:
				v = uintptr(arg.getInt())
			case argTypeInt8:
				v = uintptr(arg.getInt8())
			case argTypeInt16:
				v = uintptr(arg.getInt16())
			case argTypeInt32:
				v = uintptr(arg.getInt32())
			case argTypeInt64:
				v = uintptr(arg.getInt64())
			case argTypeUint:
				v = uintptr(arg.getUint())
			case argTypeUint8:
				v = uintptr(arg.getUint8())
			case argTypeUint16:
				v = uintptr(arg.getUint16())
			case argTypeUint32:
				v = uintptr(arg.getUint32())
			case argTypeUint64:
				v = uintptr(arg.getUint64())
			}
			switch gp {
			case 0:
				f.DI = v
			case 1:
				f.SI = v
			case 2:
				f.DX = v
			case 3:
				f.CX = v
			case 4:
				f.R8 = v
			case 5:
				f.R9 = v
			}
			gp++
		case argTypeFloat32, argTypeFloat64:
			if fp == numFP {
				// TODO: push to the stack
				// continue
				panic("not implemented")
			}
			var v uintptr
			switch arg.typ {
			case argTypeFloat32:
				v = uintptr(math.Float32bits(arg.getFloat32()))
			case argTypeFloat64:
				v = uintptr(math.Float64bits(arg.getFloat64()))
			}
			switch fp {
			case 0:
				f.X0 = v
			case 1:
				f.X1 = v
			case 2:
				f.X2 = v
			case 3:
				f.X3 = v
			case 4:
				f.X4 = v
			case 5:
				f.X5 = v
			case 6:
				f.X6 = v
			case 7:
				f.X7 = v
			}
			fp++
		}
	}
	// vararg: set %al to total number of floating point parameters in
	// vector registers.
	f.AX = fp

	libcCall(&f)

	// Keep all Go pointer arguments alive until function call completes.
	runtime.KeepAlive(args)

	switch out.typ {
	case outTypeVoid:
		return
	case outTypeUintptr:
		out.setUintptr(f.AX)
	case outTypeBool:
		out.setBool(f.AX != 0)
	case outTypeInt:
		out.setInt(int(f.AX))
	case outTypeInt8:
		out.setInt8(int8(f.AX))
	case outTypeInt16:
		out.setInt16(int16(f.AX))
	case outTypeInt32:
		out.setInt32(int32(f.AX))
	case outTypeInt64:
		out.setInt64(int64(f.AX))
	case outTypeUint:
		out.setUint(uint(f.AX))
	case outTypeUint8:
		out.setUint8(uint8(f.AX))
	case outTypeUint16:
		out.setUint16(uint16(f.AX))
	case outTypeUint32:
		out.setUint32(uint32(f.AX))
	case outTypeUint64:
		out.setUint64(uint64(f.AX))
	}
}

//go:nosplit
func libcCall(f *frame) {
	runtime_entersyscall()
	runtime_libcCall(*(*unsafe.Pointer)(unsafe.Pointer(&callABI0)), unsafe.Pointer(f))
	runtime_exitsyscall()
}

//nolint:unused // implemented in assembly
func call()

var callABI0 uintptr
