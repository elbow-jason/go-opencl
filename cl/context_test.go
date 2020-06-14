package cl

import (
	"testing"
	"unsafe"
	"errors"
	"fmt"
)

func computingCtx() (*Platform, []*Device, *Context, error) {
	platforms, err := GetPlatforms()
	if err != nil {
		return nil, nil, nil, err
	}
	if len(platforms) == 0 {
		return nil, nil, nil, errors.New("GetPlatforms had no platforms")
	}
	platform := platforms[0]
	devices, err := GetDevices(platform, DeviceTypeAll)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("GetDevice error %v", err)
	}
	context, err := CreateContext(devices)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("CreateContext error %v", err)
	}
	return platform, devices, context, nil
}

func TestCreateBufferUnsafeWorks(t *testing.T) {
	_, _, context, err := computingCtx()
	if err != nil {
		t.Fatalf("computingCtx error %v", err)
	}
	data := []byte("jason")
	memFlags := MemReadWrite | MemUseHostPtr
	ptr := unsafe.Pointer(&data[0])
	_, err = context.CreateBufferUnsafe(memFlags, len(data), ptr)
	if err != nil {
		t.Fatalf("CreateBufferUnsafe error %v", err)
	}
}