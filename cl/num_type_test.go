package cl

import (
	"testing"
	"unsafe"
)

func TestSizeofNumTypedForPrimitiveType(t *testing.T) {
	size := Int8.SizeofT()
	expected := unsafe.Sizeof(int8(0))
	if size != expected {
		t.Fatalf("Int8.SizeofT() error expected %d got %d", expected, size)
	}
	size = Uint8.SizeofT()
	expected = unsafe.Sizeof(uint8(0))
	if size != expected {
		t.Fatalf("Uint8.SizeofT() error expected %d got %d", expected, size)
	}

	size = Int16.SizeofT()
	expected = unsafe.Sizeof(int16(0))
	if size != expected {
		t.Fatalf("Int16.SizeofT() error expected %d got %d", expected, size)
	}

	size = Uint16.SizeofT()
	expected = unsafe.Sizeof(uint16(0))
	if size != expected {
		t.Fatalf("Uint16.SizeofT() error expected %d got %d", expected, size)
	}

	size = Int32.SizeofT()
	expected = unsafe.Sizeof(int32(0))
	if size != expected {
		t.Fatalf("Int32.SizeofT() error expected %d got %d", expected, size)
	}

	size = Uint32.SizeofT()
	expected = unsafe.Sizeof(uint32(0))
	if size != expected {
		t.Fatalf("Uint32.SizeofT() error expected %d got %d", expected, size)
	}

	size = Float32.SizeofT()
	expected = unsafe.Sizeof(float32(0))
	if size != expected {
		t.Fatalf("Float32.SizeofT() error expected %d got %d", expected, size)
	}

	size = Int64.SizeofT()
	expected = unsafe.Sizeof(int64(0))
	if size != expected {
		t.Fatalf("Int64.SizeofT() error expected %d got %d", expected, size)
	}

	size = Uint64.SizeofT()
	expected = unsafe.Sizeof(uint64(0))
	if size != expected {
		t.Fatalf("Uint64.SizeofT() error expected %d got %d", expected, size)
	}

	size = Float64.SizeofT()
	expected = unsafe.Sizeof(float64(0))
	if size != expected {
		t.Fatalf("Uint64.SizeofT() error expected %d got %d", expected, size)
	}

	size = Uint.SizeofT()
	expected = unsafe.Sizeof(uint(0))
	if size != expected {
		t.Fatalf("Uint.SizeofT() error expected %d got %d", expected, size)
	}
}
