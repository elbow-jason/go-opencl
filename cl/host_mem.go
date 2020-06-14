package cl

import (
	"errors"
	"unsafe"
)

// HostMem errors
var (
	ErrHostMemWasEmpty    = errors.New("HostMem slice was empty")
	ErrHostMemWasNil      = errors.New("HostMem slice was nil")
	ErrHostMemInvalidData = errors.New("HostMem data was invalid")
	ErrInvalidNumType     = errors.New("NewNumType invalid NumType")
)

// HostMem ..
type HostMem interface {
	Ptr() unsafe.Pointer
	Len() int
	NumType() NumTyped
	SizeofT() uintptr
}

// SliceF32 ..
type SliceF32 []float32

// Ptr ..
func (h SliceF32) Ptr() unsafe.Pointer {
	return unsafe.Pointer(&h[0])
}

// Len ..
func (h SliceF32) Len() int {
	return len(h)
}

// NumType ..
func (h SliceF32) NumType() NumTyped {
	return Float32
}

const sizeOfF32 = unsafe.Sizeof(float32(0))

// SizeofT ..
func (h SliceF32) SizeofT() uintptr {
	return uintptr(4)
}
