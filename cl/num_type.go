package cl

import (
	"fmt"
	"unsafe"
)

// NumTyped is the interface for Buffers and HostMem that store types of numbers.
type NumTyped interface {
	SizeofT() uintptr
	ClSrc() string
}

// PrimitiveType is a runtime enum for supported Golang number types
type PrimitiveType int

// PrimitiveType variants
const (
	Int8 PrimitiveType = iota
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Float32
	Int64
	Uint64
	Float64
	Uint
)

const (
	size8  = uintptr(1)
	size16 = uintptr(2)
	size32 = uintptr(4)
	size64 = uintptr(8)
	sizeT  = unsafe.Sizeof(uint(0))
)

func (t PrimitiveType) String() string {
	switch t {
	case Int8:
		return "Int8"
	case Uint8:
		return "Uint8"
	case Int16:
		return "Int16"
	case Uint16:
		return "Uint16"
	case Int32:
		return "Int32"
	case Uint32:
		return "Uint32"
	case Float32:
		return "Float32"
	case Int64:
		return "Int64"
	case Uint64:
		return "Uint64"
	case Float64:
		return "Float64"
	case Uint:
		return "Uint"
	default:
		panic(fmt.Sprintf("Unhandled PrimitiveType during String call"))
	}
}

// ClSrc is the OpenCL src code type string for a given PrimitiveType
func (t PrimitiveType) ClSrc() string {
	switch t {
	case Int8:
		return "char"
	case Uint8:
		return "uchar"
	case Int16:
		return "short"
	case Uint16:
		return "ushort"
	case Int32:
		return "int"
	case Uint32:
		return "uint"
	case Float32:
		return "float"
	case Int64:
		return "long"
	case Uint64:
		return "ulong"
	case Float64:
		return "double"
	case Uint:
		return "size_t"
	default:
		panic(fmt.Sprintf("Unhandled PrimitiveType during ClSrc call"))
	}
}

// SizeofT ..
func (t PrimitiveType) SizeofT() uintptr {
	switch t {
	case Int8:
		return size8
	case Uint8:
		return size8
	case Int16:
		return size16
	case Uint16:
		return size16
	case Int32:
		return size32
	case Uint32:
		return size32
	case Float32:
		return size32
	case Int64:
		return size64
	case Uint64:
		return size64
	case Float64:
		return size64
	case Uint:
		return sizeT
	default:
		panic(fmt.Sprintf("Unhandled PrimitiveType during NumType.SizeofT(): %v", t))
	}
}
