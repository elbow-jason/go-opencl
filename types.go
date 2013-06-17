package cl

// #include <OpenCL/opencl.h>
import "C"

import (
	"errors"
	"fmt"
)

var ErrUnknown = errors.New("cl: unknown error")

type ErrOther int

func (e ErrOther) Error() string {
	return fmt.Sprintf("cl: error %d", int(e))
}

var (
	ErrDeviceNotFound                     = errors.New("cl: Device Not Found")
	ErrDeviceNotAvailable                 = errors.New("cl: Device Not Available")
	ErrCompilerNotAvailable               = errors.New("cl: Compiler Not Available")
	ErrMemObjectAllocationFailure         = errors.New("cl: Mem Object Allocation Failure")
	ErrOutOfResources                     = errors.New("cl: Out Of Resources")
	ErrOutOfHostMemory                    = errors.New("cl: Out Of Host Memory")
	ErrProfilingInfoNotAvailable          = errors.New("cl: Profiling Info Not Available")
	ErrMemCopyOverlap                     = errors.New("cl: Mem Copy Overlap")
	ErrImageFormatMismatch                = errors.New("cl: Image Format Mismatch")
	ErrImageFormatNotSupported            = errors.New("cl: Image Format Not Supported")
	ErrBuildProgramFailure                = errors.New("cl: Build Program Failure")
	ErrMapFailure                         = errors.New("cl: Map Failure")
	ErrMisalignedSubBufferOffset          = errors.New("cl: Misaligned Sub Buffer Offset")
	ErrExecStatusErrorForEventsInWaitList = errors.New("cl: Exec Status Error For Events In Wait List")
	ErrCompileProgramFailure              = errors.New("cl: Compile Program Failure")
	ErrLinkerNotAvailable                 = errors.New("cl: Linker Not Available")
	ErrLinkProgramFailure                 = errors.New("cl: Link Program Failure")
	ErrDevicePartitionFailed              = errors.New("cl: Device Partition Failed")
	ErrKernelArgInfoNotAvailable          = errors.New("cl: Kernel Arg Info Not Available")
	ErrInvalidValue                       = errors.New("cl: Invalid Value")
	ErrInvalidDeviceType                  = errors.New("cl: Invalid Device Type")
	ErrInvalidPlatform                    = errors.New("cl: Invalid Platform")
	ErrInvalidDevice                      = errors.New("cl: Invalid Device")
	ErrInvalidContext                     = errors.New("cl: Invalid Context")
	ErrInvalidQueueProperties             = errors.New("cl: Invalid Queue Properties")
	ErrInvalidCommandQueue                = errors.New("cl: Invalid Command Queue")
	ErrInvalidHostPtr                     = errors.New("cl: Invalid Host Ptr")
	ErrInvalidMemObject                   = errors.New("cl: Invalid Mem Object")
	ErrInvalidImageFormatDescriptor       = errors.New("cl: Invalid Image Format Descriptor")
	ErrInvalidImageSize                   = errors.New("cl: Invalid Image Size")
	ErrInvalidSampler                     = errors.New("cl: Invalid Sampler")
	ErrInvalidBinary                      = errors.New("cl: Invalid Binary")
	ErrInvalidBuildOptions                = errors.New("cl: Invalid Build Options")
	ErrInvalidProgram                     = errors.New("cl: Invalid Program")
	ErrInvalidProgramExecutable           = errors.New("cl: Invalid Program Executable")
	ErrInvalidKernelName                  = errors.New("cl: Invalid Kernel Name")
	ErrInvalidKernelDefinition            = errors.New("cl: Invalid Kernel Definition")
	ErrInvalidKernel                      = errors.New("cl: Invalid Kernel")
	ErrInvalidArgIndex                    = errors.New("cl: Invalid Arg Index")
	ErrInvalidArgValue                    = errors.New("cl: Invalid Arg Value")
	ErrInvalidArgSize                     = errors.New("cl: Invalid Arg Size")
	ErrInvalidKernelArgs                  = errors.New("cl: Invalid Kernel Args")
	ErrInvalidWorkDimension               = errors.New("cl: Invalid Work Dimension")
	ErrInvalidWorkGroupSize               = errors.New("cl: Invalid Work Group Size")
	ErrInvalidWorkItemSize                = errors.New("cl: Invalid Work Item Size")
	ErrInvalidGlobalOffset                = errors.New("cl: Invalid Global Offset")
	ErrInvalidEventWaitList               = errors.New("cl: Invalid Event Wait List")
	ErrInvalidEvent                       = errors.New("cl: Invalid Event")
	ErrInvalidOperation                   = errors.New("cl: Invalid Operation")
	ErrInvalidGlObject                    = errors.New("cl: Invalid Gl Object")
	ErrInvalidBufferSize                  = errors.New("cl: Invalid Buffer Size")
	ErrInvalidMipLevel                    = errors.New("cl: Invalid Mip Level")
	ErrInvalidGlobalWorkSize              = errors.New("cl: Invalid Global Work Size")
	ErrInvalidProperty                    = errors.New("cl: Invalid Property")
	ErrInvalidImageDescriptor             = errors.New("cl: Invalid Image Descriptor")
	ErrInvalidCompilerOptions             = errors.New("cl: Invalid Compiler Options")
	ErrInvalidLinkerOptions               = errors.New("cl: Invalid Linker Options")
	ErrInvalidDevicePartitionCount        = errors.New("cl: Invalid Device Partition Count")
)
var errorMap = map[C.cl_int]error{
	C.CL_SUCCESS:                                   nil,
	C.CL_DEVICE_NOT_FOUND:                          ErrDeviceNotFound,
	C.CL_DEVICE_NOT_AVAILABLE:                      ErrDeviceNotAvailable,
	C.CL_COMPILER_NOT_AVAILABLE:                    ErrCompilerNotAvailable,
	C.CL_MEM_OBJECT_ALLOCATION_FAILURE:             ErrMemObjectAllocationFailure,
	C.CL_OUT_OF_RESOURCES:                          ErrOutOfResources,
	C.CL_OUT_OF_HOST_MEMORY:                        ErrOutOfHostMemory,
	C.CL_PROFILING_INFO_NOT_AVAILABLE:              ErrProfilingInfoNotAvailable,
	C.CL_MEM_COPY_OVERLAP:                          ErrMemCopyOverlap,
	C.CL_IMAGE_FORMAT_MISMATCH:                     ErrImageFormatMismatch,
	C.CL_IMAGE_FORMAT_NOT_SUPPORTED:                ErrImageFormatNotSupported,
	C.CL_BUILD_PROGRAM_FAILURE:                     ErrBuildProgramFailure,
	C.CL_MAP_FAILURE:                               ErrMapFailure,
	C.CL_MISALIGNED_SUB_BUFFER_OFFSET:              ErrMisalignedSubBufferOffset,
	C.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST: ErrExecStatusErrorForEventsInWaitList,
	C.CL_COMPILE_PROGRAM_FAILURE:                   ErrCompileProgramFailure,
	C.CL_LINKER_NOT_AVAILABLE:                      ErrLinkerNotAvailable,
	C.CL_LINK_PROGRAM_FAILURE:                      ErrLinkProgramFailure,
	C.CL_DEVICE_PARTITION_FAILED:                   ErrDevicePartitionFailed,
	C.CL_KERNEL_ARG_INFO_NOT_AVAILABLE:             ErrKernelArgInfoNotAvailable,
	C.CL_INVALID_VALUE:                             ErrInvalidValue,
	C.CL_INVALID_DEVICE_TYPE:                       ErrInvalidDeviceType,
	C.CL_INVALID_PLATFORM:                          ErrInvalidPlatform,
	C.CL_INVALID_DEVICE:                            ErrInvalidDevice,
	C.CL_INVALID_CONTEXT:                           ErrInvalidContext,
	C.CL_INVALID_QUEUE_PROPERTIES:                  ErrInvalidQueueProperties,
	C.CL_INVALID_COMMAND_QUEUE:                     ErrInvalidCommandQueue,
	C.CL_INVALID_HOST_PTR:                          ErrInvalidHostPtr,
	C.CL_INVALID_MEM_OBJECT:                        ErrInvalidMemObject,
	C.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR:           ErrInvalidImageFormatDescriptor,
	C.CL_INVALID_IMAGE_SIZE:                        ErrInvalidImageSize,
	C.CL_INVALID_SAMPLER:                           ErrInvalidSampler,
	C.CL_INVALID_BINARY:                            ErrInvalidBinary,
	C.CL_INVALID_BUILD_OPTIONS:                     ErrInvalidBuildOptions,
	C.CL_INVALID_PROGRAM:                           ErrInvalidProgram,
	C.CL_INVALID_PROGRAM_EXECUTABLE:                ErrInvalidProgramExecutable,
	C.CL_INVALID_KERNEL_NAME:                       ErrInvalidKernelName,
	C.CL_INVALID_KERNEL_DEFINITION:                 ErrInvalidKernelDefinition,
	C.CL_INVALID_KERNEL:                            ErrInvalidKernel,
	C.CL_INVALID_ARG_INDEX:                         ErrInvalidArgIndex,
	C.CL_INVALID_ARG_VALUE:                         ErrInvalidArgValue,
	C.CL_INVALID_ARG_SIZE:                          ErrInvalidArgSize,
	C.CL_INVALID_KERNEL_ARGS:                       ErrInvalidKernelArgs,
	C.CL_INVALID_WORK_DIMENSION:                    ErrInvalidWorkDimension,
	C.CL_INVALID_WORK_GROUP_SIZE:                   ErrInvalidWorkGroupSize,
	C.CL_INVALID_WORK_ITEM_SIZE:                    ErrInvalidWorkItemSize,
	C.CL_INVALID_GLOBAL_OFFSET:                     ErrInvalidGlobalOffset,
	C.CL_INVALID_EVENT_WAIT_LIST:                   ErrInvalidEventWaitList,
	C.CL_INVALID_EVENT:                             ErrInvalidEvent,
	C.CL_INVALID_OPERATION:                         ErrInvalidOperation,
	C.CL_INVALID_GL_OBJECT:                         ErrInvalidGlObject,
	C.CL_INVALID_BUFFER_SIZE:                       ErrInvalidBufferSize,
	C.CL_INVALID_MIP_LEVEL:                         ErrInvalidMipLevel,
	C.CL_INVALID_GLOBAL_WORK_SIZE:                  ErrInvalidGlobalWorkSize,
	C.CL_INVALID_PROPERTY:                          ErrInvalidProperty,
	C.CL_INVALID_IMAGE_DESCRIPTOR:                  ErrInvalidImageDescriptor,
	C.CL_INVALID_COMPILER_OPTIONS:                  ErrInvalidCompilerOptions,
	C.CL_INVALID_LINKER_OPTIONS:                    ErrInvalidLinkerOptions,
	C.CL_INVALID_DEVICE_PARTITION_COUNT:            ErrInvalidDevicePartitionCount,
}

func toError(code C.cl_int) error {
	if err, ok := errorMap[code]; ok {
		return err
	}
	return ErrOther(code)
}

type MemFlag int

const (
	MemReadWrite     MemFlag = C.CL_MEM_READ_WRITE
	MemWriteOnly     MemFlag = C.CL_MEM_WRITE_ONLY
	MemReadOnly      MemFlag = C.CL_MEM_READ_ONLY
	MemUseHostPtr    MemFlag = C.CL_MEM_USE_HOST_PTR
	MemAllocHostPtr  MemFlag = C.CL_MEM_ALLOC_HOST_PTR
	MemCopyHostPtr   MemFlag = C.CL_MEM_COPY_HOST_PTR
	MemHostWriteOnly MemFlag = C.CL_MEM_HOST_WRITE_ONLY // OpenCL 1.2
	MemHostReadOnly  MemFlag = C.CL_MEM_HOST_READ_ONLY  // OpenCL 1.2
	MemHostNoAccess  MemFlag = C.CL_MEM_HOST_NO_ACCESS  // OpenCL 1.2
)

type MemObjectType int

const (
	MemObjectTypeBuffer        MemObjectType = C.CL_MEM_OBJECT_BUFFER
	MemObjectTypeImage2D       MemObjectType = C.CL_MEM_OBJECT_IMAGE2D
	MemObjectTypeImage3D       MemObjectType = C.CL_MEM_OBJECT_IMAGE3D
	MemObjectTypeImage2DArray  MemObjectType = C.CL_MEM_OBJECT_IMAGE2D_ARRAY
	MemObjectTypeImage1D       MemObjectType = C.CL_MEM_OBJECT_IMAGE1D
	MemObjectTypeImage1DArray  MemObjectType = C.CL_MEM_OBJECT_IMAGE1D_ARRAY
	MemObjectTypeImage1DBuffer MemObjectType = C.CL_MEM_OBJECT_IMAGE1D_BUFFER
)

type ChannelOrder int

const (
	ChannelOrderR            ChannelOrder = C.CL_R
	ChannelOrderA            ChannelOrder = C.CL_A
	ChannelOrderRG           ChannelOrder = C.CL_RG
	ChannelOrderRA           ChannelOrder = C.CL_RA
	ChannelOrderRGB          ChannelOrder = C.CL_RGB
	ChannelOrderRGBA         ChannelOrder = C.CL_RGBA
	ChannelOrderBGRA         ChannelOrder = C.CL_BGRA
	ChannelOrderARGB         ChannelOrder = C.CL_ARGB
	ChannelOrderIntensity    ChannelOrder = C.CL_INTENSITY
	ChannelOrderLuminance    ChannelOrder = C.CL_LUMINANCE
	ChannelOrderRx           ChannelOrder = C.CL_Rx
	ChannelOrderRGx          ChannelOrder = C.CL_RGx
	ChannelOrderRGBx         ChannelOrder = C.CL_RGBx
	ChannelOrderDepth        ChannelOrder = C.CL_DEPTH
	ChannelOrderDepthStencil ChannelOrder = C.CL_DEPTH_STENCIL
)

var channelOrderNameMap = map[ChannelOrder]string{
	ChannelOrderR:            "R",
	ChannelOrderA:            "A",
	ChannelOrderRG:           "RG",
	ChannelOrderRA:           "RA",
	ChannelOrderRGB:          "RGB",
	ChannelOrderRGBA:         "RGBA",
	ChannelOrderBGRA:         "BGRA",
	ChannelOrderARGB:         "ARGB",
	ChannelOrderIntensity:    "Intensity",
	ChannelOrderLuminance:    "Luminance",
	ChannelOrderRx:           "Rx",
	ChannelOrderRGx:          "RGx",
	ChannelOrderRGBx:         "RGBx",
	ChannelOrderDepth:        "Depth",
	ChannelOrderDepthStencil: "DepthStencil",
}

func (co ChannelOrder) String() string {
	name := channelOrderNameMap[co]
	if name == "" {
		name = "Unknown"
	}
	return name
}

type ChannelDataType int

const (
	ChannelDataTypeSNormInt8      ChannelDataType = C.CL_SNORM_INT8
	ChannelDataTypeSNormInt16     ChannelDataType = C.CL_SNORM_INT16
	ChannelDataTypeUNormInt8      ChannelDataType = C.CL_UNORM_INT8
	ChannelDataTypeUNormInt16     ChannelDataType = C.CL_UNORM_INT16
	ChannelDataTypeUNormShort565  ChannelDataType = C.CL_UNORM_SHORT_565
	ChannelDataTypeUNormShort555  ChannelDataType = C.CL_UNORM_SHORT_555
	ChannelDataTypeUNormInt101010 ChannelDataType = C.CL_UNORM_INT_101010
	ChannelDataTypeSignedInt8     ChannelDataType = C.CL_SIGNED_INT8
	ChannelDataTypeSignedInt16    ChannelDataType = C.CL_SIGNED_INT16
	ChannelDataTypeSignedInt32    ChannelDataType = C.CL_SIGNED_INT32
	ChannelDataTypeUnsignedInt8   ChannelDataType = C.CL_UNSIGNED_INT8
	ChannelDataTypeUnsignedInt16  ChannelDataType = C.CL_UNSIGNED_INT16
	ChannelDataTypeUnsignedInt32  ChannelDataType = C.CL_UNSIGNED_INT32
	ChannelDataTypeHalfFloat      ChannelDataType = C.CL_HALF_FLOAT
	ChannelDataTypeFloat          ChannelDataType = C.CL_FLOAT
	ChannelDataTypeUNormInt24     ChannelDataType = C.CL_UNORM_INT24
)

var channelDataTypeNameMap = map[ChannelDataType]string{
	ChannelDataTypeSNormInt8:      "SNormInt8",
	ChannelDataTypeSNormInt16:     "SNormInt16",
	ChannelDataTypeUNormInt8:      "UNormInt8",
	ChannelDataTypeUNormInt16:     "UNormInt16",
	ChannelDataTypeUNormShort565:  "UNormShort565",
	ChannelDataTypeUNormShort555:  "UNormShort555",
	ChannelDataTypeUNormInt101010: "UNormInt101010",
	ChannelDataTypeSignedInt8:     "SignedInt8",
	ChannelDataTypeSignedInt16:    "SignedInt16",
	ChannelDataTypeSignedInt32:    "SignedInt32",
	ChannelDataTypeUnsignedInt8:   "UnsignedInt8",
	ChannelDataTypeUnsignedInt16:  "UnsignedInt16",
	ChannelDataTypeUnsignedInt32:  "UnsignedInt32",
	ChannelDataTypeHalfFloat:      "HalfFloat",
	ChannelDataTypeFloat:          "Float",
	ChannelDataTypeUNormInt24:     "UNormInt24",
}

func (ct ChannelDataType) String() string {
	name := channelDataTypeNameMap[ct]
	if name == "" {
		name = "Unknown"
	}
	return name
}

type ImageFormat struct {
	ChannelOrder    ChannelOrder
	ChannelDataType ChannelDataType
}

func (f ImageFormat) toCl() C.cl_image_format {
	var format C.cl_image_format
	format.image_channel_order = C.cl_channel_order(f.ChannelOrder)
	format.image_channel_data_type = C.cl_channel_type(f.ChannelDataType)
	return format
}

type ImageDescription struct {
	Type                            MemObjectType
	Width, Height, Depth            int
	ArraySize, RowPitch, SlicePitch int
	NumMipLevels, NumSamples        int
	Buffer                          *Buffer
}

func (d ImageDescription) toCl() C.cl_image_desc {
	var desc C.cl_image_desc
	desc.image_type = C.cl_mem_object_type(d.Type)
	desc.image_width = C.size_t(d.Width)
	desc.image_height = C.size_t(d.Height)
	desc.image_depth = C.size_t(d.Depth)
	desc.image_array_size = C.size_t(d.ArraySize)
	desc.image_row_pitch = C.size_t(d.RowPitch)
	desc.image_slice_pitch = C.size_t(d.SlicePitch)
	desc.num_mip_levels = C.cl_uint(d.NumMipLevels)
	desc.num_samples = C.cl_uint(d.NumSamples)
	desc.buffer = nil
	if d.Buffer != nil {
		desc.buffer = d.Buffer.clBuffer
	}
	return desc
}

type Event C.cl_event

func clBool(b bool) C.cl_bool {
	if b {
		return C.CL_TRUE
	}
	return C.CL_FALSE
}

func sizeT3(i3 [3]int) [3]C.size_t {
	var val [3]C.size_t
	val[0] = C.size_t(i3[0])
	val[1] = C.size_t(i3[1])
	val[2] = C.size_t(i3[2])
	return val
}

func eventListPtr(el []Event) *C.cl_event {
	if el == nil {
		return nil
	}
	return (*C.cl_event)(&el[0])
}