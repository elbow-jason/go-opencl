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
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// ProgramBuildError ..
type ProgramBuildError string

// Error ..
func (e ProgramBuildError) Error() string {
	return fmt.Sprintf("cl: build error (%s)", string(e))
}

// Program is the cl_program wrapping struct
type Program struct {
	clProgram C.cl_program
	devices   []*Device
}

func releaseProgram(p *Program) {
	if p.clProgram != nil {
		C.clReleaseProgram(p.clProgram)
		p.clProgram = nil
	}
}

// Release decrements the OpenCL atomic reference count of the underlying program pointer.
func (p *Program) Release() {
	releaseProgram(p)
}

// BuildProgram compiles the source code of the program on the given devices.
func (p *Program) BuildProgram(devices []*Device, options string) error {
	var cOptions *C.char
	if options != "" {
		cOptions = C.CString(options)
		defer C.free(unsafe.Pointer(cOptions))
	}
	var deviceList []C.cl_device_id
	var deviceListPtr *C.cl_device_id
	numDevices := C.cl_uint(0)
	if devices != nil && len(devices) > 0 {
		deviceList = buildDeviceIDList(devices)
		deviceListPtr = &deviceList[0]
	}
	statusCode := C.clBuildProgram(p.clProgram, numDevices, deviceListPtr, cOptions, nil, nil)
	err := toError(statusCode)
	if err != nil {
		return err
	}
	return nil
}

// BuildLogs ..
func (p Program) BuildLogs() ([]string, error) {
	logs := make([]string, len(p.devices))
	var bLen C.size_t
	var err error
	for _, device := range p.devices {
		// Get the byteSize
		err = toError(C.clGetProgramBuildInfo(p.clProgram, device.id, C.CL_PROGRAM_BUILD_LOG, 0, nil, &bLen))
		if err != nil {
			return nil, err
		}
		buffer := make([]byte, bLen)
		err = toError(C.clGetProgramBuildInfo(p.clProgram, device.id, C.CL_PROGRAM_BUILD_LOG, bLen, unsafe.Pointer(&buffer[0]), nil))
		if err != nil {
			return nil, err
		}
		logs = append(logs, string(buffer))
	}
	return logs, nil
}

// CreateKernel returns the *Kernel of the given name.
func (p *Program) CreateKernel(name string) (*Kernel, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var err C.cl_int
	clKernel := C.clCreateKernel(p.clProgram, cName, &err)
	if err != C.CL_SUCCESS {
		return nil, toError(err)
	}
	kernel := &Kernel{clKernel: clKernel, name: name}
	runtime.SetFinalizer(kernel, releaseKernel)
	return kernel, nil
}
