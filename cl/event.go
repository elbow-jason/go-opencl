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
	"runtime"
)


// Event is the cl_event wrapping struct
type Event struct {
	clEvent C.cl_event
}

func releaseEvent(ev *Event) {
	if ev.clEvent != nil {
		C.clReleaseEvent(ev.clEvent)
		ev.clEvent = nil
	}
}

// Release decrements the OpenCL atomic reference count for the underlying cl_event.
func (e *Event) Release() {
	releaseEvent(e)
}

// GetEventProfilingInfo returns the profiliing value for the given ProfilingInfo.
// This info can be used to tune/benchmark execution.
func (e *Event) GetEventProfilingInfo(paramName ProfilingInfo) (int64, error) {
	var paramValue C.cl_ulong
	if err := C.clGetEventProfilingInfo(e.clEvent, C.cl_profiling_info(paramName), C.size_t(unsafe.Sizeof(paramValue)), unsafe.Pointer(&paramValue), nil); err != C.CL_SUCCESS {
		return 0, toError(err)
	}
	return int64(paramValue), nil
}

// SetUserEventStatus sets the execution status of a user event object.
//
// `status` specifies the new execution status to be set and
// can be CL_COMPLETE or a negative integer value to indicate
// an error. A negative integer value causes all enqueued commands
// that wait on this user event to be terminated. clSetUserEventStatus
// can only be called once to change the execution status of event.
func (e *Event) SetUserEventStatus(status int) error {
	return toError(C.clSetUserEventStatus(e.clEvent, C.cl_int(status)))
}

// WaitForEvents waits on the host thread for commands identified by event objects in
// events to complete. A command is considered complete if its execution
// status is CL_COMPLETE or a negative value. The events specified in
// event_list act as synchronization points.
//
// If the cl_khr_gl_event extension is enabled, event objects can also be
// used to reflect the status of an OpenGL sync object. The sync object
// in turn refers to a fence command executing in an OpenGL command
// stream. This provides another method of coordinating sharing of buffers
// and images between OpenGL and OpenCL.
func WaitForEvents(events []*Event) error {
	return toError(C.clWaitForEvents(C.cl_uint(len(events)), eventListPtr(events)))
}

func newEvent(clEvent C.cl_event) *Event {
	ev := &Event{clEvent: clEvent}
	runtime.SetFinalizer(ev, releaseEvent)
	return ev
}

func eventListPtr(el []*Event) *C.cl_event {
	if el == nil {
		return nil
	}
	elist := make([]C.cl_event, len(el))
	for i, e := range el {
		elist[i] = e.clEvent
	}
	return (*C.cl_event)(&elist[0])
}

// NewWaitlist turns variadric *cl.Event args into nil (if empty) or []*cl.Event (which is a Waitlist)
func NewWaitlist(args ...*Event) []*Event {
	if len(args) == 0 {
		return nil
	}
	return args
}