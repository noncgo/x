//go:build darwin && amd64
// +build darwin,amd64

package cabi

import (
	"math"
	"unsafe"
)

// Compile-time check to ensure that we are at most on a 64-bit platform.
var _ [8 - unsafe.Sizeof(uintptr(0))]struct{}

type argType uint8

const (
	argTypeVoid argType = iota
	argTypeUnsafePointer
	argTypeString
	argTypeBytes
	argTypeUintptr
	argTypeBool
	argTypeInt
	argTypeInt8
	argTypeInt16
	argTypeInt32
	argTypeInt64
	argTypeUint
	argTypeUint8
	argTypeUint16
	argTypeUint32
	argTypeUint64
	argTypeFloat32
	argTypeFloat64
)

// Arg is a function call argument value.
type Arg struct {
	val uint64
	str string
	buf []byte
	ptr unsafe.Pointer
	typ argType
}

// UnsafePointer returns a function call argument value for unsafe.Pointer type.
func UnsafePointer(v unsafe.Pointer) Arg {
	return Arg{
		typ: argTypeUnsafePointer,
		ptr: v,
	}
}

// String returns a function call argument value for string type.
func String(v string) Arg {
	return Arg{
		typ: argTypeString,
		str: v,
	}
}

// Bytes returns a function call argument value for []byte type.
func Bytes(v []byte) Arg {
	return Arg{
		typ: argTypeBytes,
		buf: v,
	}
}

// Uintptr returns a function call argument value for uintptr type.
func Uintptr(v uintptr) Arg {
	return Arg{
		typ: argTypeUintptr,
		val: uint64(v),
	}
}

// Bool returns a function call argument value for bool type.
func Bool(v bool) Arg {
	var val uint64
	if v {
		val = 1
	}
	return Arg{
		typ: argTypeBool,
		val: val,
	}
}

// Int returns a function call argument value for int type.
func Int(v int) Arg {
	return Arg{
		typ: argTypeInt,
		val: uint64(v),
	}
}

// Int8 returns a function call argument value for int8 type.
func Int8(v int8) Arg {
	return Arg{
		typ: argTypeInt8,
		val: uint64(v),
	}
}

// Int16 returns a function call argument value for int32 type.
func Int16(v int16) Arg {
	return Arg{
		typ: argTypeInt16,
		val: uint64(v),
	}
}

// Int32 returns a function call argument value for int32 type.
func Int32(v int32) Arg {
	return Arg{
		typ: argTypeInt32,
		val: uint64(v),
	}
}

// Int64 returns a function call argument value for int64 type.
func Int64(v int64) Arg {
	return Arg{
		typ: argTypeInt64,
		val: uint64(v),
	}
}

// Uint returns a function call argument value for uint type.
func Uint(v uint) Arg {
	return Arg{
		typ: argTypeUint,
		val: uint64(v),
	}
}

// Uint8 returns a function call argument value for uint8 type.
func Uint8(v uint8) Arg {
	return Arg{
		typ: argTypeUint8,
		val: uint64(v),
	}
}

// Uint16 returns a function call argument value for uint32 type.
func Uint16(v uint16) Arg {
	return Arg{
		typ: argTypeUint16,
		val: uint64(v),
	}
}

// Uint32 returns a function call argument value for uint32 type.
func Uint32(v uint32) Arg {
	return Arg{
		typ: argTypeUint32,
		val: uint64(v),
	}
}

// Uint64 returns a function call argument value for uint64 type.
func Uint64(v uint64) Arg {
	return Arg{
		typ: argTypeUint64,
		val: v,
	}
}

// Float32 returns a function call argument value for float32 type.
func Float32(v float32) Arg {
	return Arg{
		typ: argTypeFloat32,
		val: uint64(math.Float32bits(v)),
	}
}

// Float64 returns a function call argument value for float64 type.
func Float64(v float64) Arg {
	return Arg{
		typ: argTypeFloat64,
		val: math.Float64bits(v),
	}
}

func (a *Arg) getUnsafePointer() unsafe.Pointer {
	return a.ptr
}

func (a *Arg) getString() string {
	return a.str
}

func (a *Arg) getBytes() []byte {
	return a.buf
}

func (a *Arg) getUintptr() uintptr {
	return uintptr(a.val)
}

func (a *Arg) getBool() bool {
	return a.val != 0
}

func (a *Arg) getInt() int {
	return int(a.val)
}

func (a *Arg) getInt8() int8 {
	return int8(a.val)
}

func (a *Arg) getInt16() int16 {
	return int16(a.val)
}

func (a *Arg) getInt32() int32 {
	return int32(a.val)
}

func (a *Arg) getInt64() int64 {
	return int64(a.val)
}

func (a *Arg) getUint() uint {
	return uint(a.val)
}

func (a *Arg) getUint8() uint8 {
	return uint8(a.val)
}

func (a *Arg) getUint16() uint16 {
	return uint16(a.val)
}

func (a *Arg) getUint32() uint32 {
	return uint32(a.val)
}

func (a *Arg) getUint64() uint64 {
	return a.val
}

func (a *Arg) getFloat32() float32 {
	return math.Float32frombits(uint32(a.val))
}

func (a *Arg) getFloat64() float64 {
	return math.Float64frombits(a.val)
}
