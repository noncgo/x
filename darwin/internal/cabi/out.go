//go:build darwin && amd64
// +build darwin,amd64

package cabi

import (
	"unsafe"
)

type outType uint8

const (
	outTypeVoid outType = iota
	outTypeUintptr
	outTypeBool
	outTypeInt
	outTypeInt8
	outTypeInt16
	outTypeInt32
	outTypeInt64
	outTypeUint
	outTypeUint8
	outTypeUint16
	outTypeUint32
	outTypeUint64
)

// Out is a function call output value.
type Out struct {
	typ outType
	val unsafe.Pointer
}

// Void returns a function call output value for void type.
func Void() Out {
	return Out{
		typ: outTypeVoid,
	}
}

// OutUintptr returns a function call output value for uintptr type.
func OutUintptr(p *uintptr) Out {
	return Out{
		typ: outTypeUintptr,
		val: unsafe.Pointer(p),
	}
}

// OutBool returns a function call output value for bool type.
func OutBool(p *bool) Out {
	return Out{
		typ: outTypeBool,
		val: unsafe.Pointer(p),
	}
}

// OutInt returns a function call output value for int type.
func OutInt(p *int) Out {
	return Out{
		typ: outTypeInt,
		val: unsafe.Pointer(p),
	}
}

// OutInt8 returns a function call output value for int8 type.
func OutInt8(p *int8) Out {
	return Out{
		typ: outTypeInt8,
		val: unsafe.Pointer(p),
	}
}

// OutInt16 returns a function call output value for int16 type.
func OutInt16(p *int16) Out {
	return Out{
		typ: outTypeInt16,
		val: unsafe.Pointer(p),
	}
}

// OutInt32 returns a function call output value for int32 type.
func OutInt32(p *int32) Out {
	return Out{
		typ: outTypeInt32,
		val: unsafe.Pointer(p),
	}
}

// OutInt64 returns a function call output value for int64 type.
func OutInt64(p *int64) Out {
	return Out{
		typ: outTypeInt64,
		val: unsafe.Pointer(p),
	}
}

// OutUint returns a function call output value for uint type.
func OutUint(p *uint) Out {
	return Out{
		typ: outTypeUint,
		val: unsafe.Pointer(p),
	}
}

// OutUint8 returns a function call output value for uint8 type.
func OutUint8(p *uint8) Out {
	return Out{
		typ: outTypeUint8,
		val: unsafe.Pointer(p),
	}
}

// OutUint16 returns a function call output value for uint16 type.
func OutUint16(p *uint16) Out {
	return Out{
		typ: outTypeUint16,
		val: unsafe.Pointer(p),
	}
}

// OutUint32 returns a function call output value for uint32 type.
func OutUint32(p *uint32) Out {
	return Out{
		typ: outTypeUint32,
		val: unsafe.Pointer(p),
	}
}

// OutUint64 returns a function call output value for uint64 type.
func OutUint64(p *uint64) Out {
	return Out{
		typ: outTypeUint64,
		val: unsafe.Pointer(p),
	}
}

func (o *Out) setUintptr(v uintptr) {
	*(*uintptr)(o.val) = v
}

func (o *Out) setBool(v bool) {
	*(*bool)(o.val) = v
}

func (o *Out) setInt(v int) {
	*(*int)(o.val) = v
}

func (o *Out) setInt8(v int8) {
	*(*int8)(o.val) = v
}

func (o *Out) setInt16(v int16) {
	*(*int16)(o.val) = v
}

func (o *Out) setInt32(v int32) {
	*(*int32)(o.val) = v
}

func (o *Out) setInt64(v int64) {
	*(*int64)(o.val) = v
}

func (o *Out) setUint(v uint) {
	*(*uint)(o.val) = v
}

func (o *Out) setUint8(v uint8) {
	*(*uint8)(o.val) = v
}

func (o *Out) setUint16(v uint16) {
	*(*uint16)(o.val) = v
}

func (o *Out) setUint32(v uint32) {
	*(*uint32)(o.val) = v
}

func (o *Out) setUint64(v uint64) {
	*(*uint64)(o.val) = v
}
