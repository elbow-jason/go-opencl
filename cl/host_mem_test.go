package cl

import (
	"testing"
	"unsafe"
)

func takesHostMem(h HostMem) HostMem {
	return h
}

func TestHostMemForSliceFloat32(t *testing.T) {
	slice := []float32{1.0, 2.0, 3.0}
	sliceF32 := SliceF32(slice)
	hostMem := takesHostMem(sliceF32)
	sizeofT := hostMem.SizeofT()
	expected := unsafe.Sizeof(float32(0))
	if sizeofT != expected {
		t.Fatalf("hostMem.SizeofT() was incorrect expected %d got %d", expected, sizeofT)
	}
	length := hostMem.Len()
	if length != 3 {
		t.Fatalf("hostMem.Len() was incorrect expected 3 got %d", length)
	}
	ptr := hostMem.Ptr()
	addr := unsafe.Pointer(&slice[0])
	if ptr != addr {
		t.Fatalf("hostMem.Ptr() was incorrect address %v got %v", ptr, addr)
	}
	numType := hostMem.NumType()
	if numType.(PrimitiveType) != Float32 {
		t.Fatalf("SliceF32 was not NumTyped as PrimitiveType Float32")
	}
}
