package cl

/*
#cgo CFLAGS: -I CL -w
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL
#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// ErrUnsupportedArgumentType ..
type ErrUnsupportedArgumentType struct {
	Index int
	Value interface{}
}

func (e ErrUnsupportedArgumentType) Error() string {
	return fmt.Sprintf("cl: unsupported argument type for index %d: %+v", e.Index, e.Value)
}

// Kernel ..
type Kernel struct {
	clKernel C.cl_kernel
	name     string
}

// LocalBuffer ..
type LocalBuffer int

func releaseKernel(k *Kernel) {
	if k.clKernel != nil {
		C.clReleaseKernel(k.clKernel)
		k.clKernel = nil
	}
}

// Release ..
func (k *Kernel) Release() {
	releaseKernel(k)
}

// SetArgs ..
func (k *Kernel) SetArgs(args ...interface{}) error {
	for index, arg := range args {
		if err := k.SetArg(index, arg); err != nil {
			return err
		}
	}
	return nil
}

// SetArg sets the given arg at the given index on the Kernel
func (k *Kernel) SetArg(index int, arg interface{}) error {
	switch val := arg.(type) {
	case *MemObject:
		return k.SetArgBuffer(index, val)
	case LocalBuffer:
		return k.SetArgLocal(index, int(val))
	default:
		return k.SetArgNumber(index, arg)
	}
}

// SetArgBuffer ..
func (k *Kernel) SetArgBuffer(index int, buffer *MemObject) error {
	return k.SetArgUnsafe(index, int(unsafe.Sizeof(buffer.clMem)), unsafe.Pointer(&buffer.clMem))
}

// SetArgLocal ..
func (k *Kernel) SetArgLocal(index int, size int) error {
	return k.SetArgUnsafe(index, size, nil)
}

// SetArgNumber ..
func (k *Kernel) SetArgNumber(index int, arg interface{}) error {
	switch val := arg.(type) {
	case uint8, int8, uint16, int16, uint32, int32, float32, int64, uint64, float64:
		return k.SetArgUnsafe(index, int(unsafe.Sizeof(val)), unsafe.Pointer(&val))
	default:
		return ErrUnsupportedArgumentType{Index: index, Value: arg}
	}
}

// SetArgUnsafe ..
func (k *Kernel) SetArgUnsafe(index, argSize int, arg unsafe.Pointer) error {
	return toError(C.clSetKernelArg(k.clKernel, C.cl_uint(index), C.size_t(argSize), arg))
}

// PreferredWorkGroupSizeMultiple ..
func (k *Kernel) PreferredWorkGroupSizeMultiple(device *Device) (int, error) {
	var size C.size_t
	err := C.clGetKernelWorkGroupInfo(k.clKernel, device.nullableId(), C.CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE, C.size_t(unsafe.Sizeof(size)), unsafe.Pointer(&size), nil)
	return int(size), toError(err)
}

// WorkGroupSize ..
func (k *Kernel) WorkGroupSize(device *Device) (int, error) {
	var size C.size_t
	err := C.clGetKernelWorkGroupInfo(k.clKernel, device.nullableId(), C.CL_KERNEL_WORK_GROUP_SIZE, C.size_t(unsafe.Sizeof(size)), unsafe.Pointer(&size), nil)
	return int(size), toError(err)
}

func (k *Kernel) NumArgs() (int, error) {
	var num C.cl_uint
	err := C.clGetKernelInfo(k.clKernel, C.CL_KERNEL_NUM_ARGS, C.size_t(unsafe.Sizeof(num)), unsafe.Pointer(&num), nil)
	return int(num), toError(err)
}
