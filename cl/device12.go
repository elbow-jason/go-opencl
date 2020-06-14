// +build !cl10

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
	"unsafe"
	"strings"
)

// FPConfigCorrectlyRoundedDivideSqrt ..
const FPConfigCorrectlyRoundedDivideSqrt FPConfig = C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT

func init() {
	fpConfigNameMap[FPConfigCorrectlyRoundedDivideSqrt] = "CorrectlyRoundedDivideSqrt"
}

// BuiltInKernels .. 
func (d *Device) BuiltInKernels() []string {
	// From specification:
	// A semi-colon separated list of built-in kernels supported by the device. An
	// empty string is returned if no built-in kernels are supported by the
	// device.
	str, _ := d.getInfoString(C.CL_DEVICE_BUILT_IN_KERNELS, true)
	return strings.Split(str, ";")
}

// LinkerAvailable is false if the implementation does not have a linker
// available and is true if the linker is available. This can be false for the
// embedded platform profile only. This must be true if
// device.CompilerAvailable() is true.
func (d *Device) LinkerAvailable() bool {
	val, _ := d.getInfoBool(C.CL_DEVICE_LINKER_AVAILABLE, true)
	return val
}

// ParentDevice ..
func (d *Device) ParentDevice() *Device {
	var deviceID C.cl_device_id
	if err := C.clGetDeviceInfo(d.id, C.CL_DEVICE_PARENT_DEVICE, C.size_t(unsafe.Sizeof(deviceID)), unsafe.Pointer(&deviceID), nil); err != C.CL_SUCCESS {
		panic("ParentDevice failed")
	}
	if deviceID == nil {
		return nil
	}
	return &Device{id: deviceID}
}

// ImageMaxBufferSize is the max number of pixels for a 1D image created from a
// buffer object. The minimum value is 65536 if CL_DEVICE_IMAGE_SUPPORT is
// CL_TRUE.
func (d *Device) ImageMaxBufferSize() int {
	val, _ := d.getInfoSize(C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE, true)
	return int(val)
}

// ImageMaxArraySize is the max number of images in a 1D or 2D image array. The
// minimum value is 2048 if CL_DEVICE_IMAGE_SUPPORT is CL_TRUE.
func (d *Device) ImageMaxArraySize() int {
	val, _ := d.getInfoSize(C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE, true)
	return int(val)
}
