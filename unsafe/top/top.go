package top

import (
	top64 "github.com/gomacro/top/64/top"
	top32 "github.com/gomacro/top/32/top"
	top8 "github.com/gomacro/top/8/top"
	"reflect"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////
func elemsize(slice interface{}) uintptr {
	return uintptr(reflect.TypeOf(slice).Elem().Size())
}
func mvetype(dst, src *interface{}) {
	*(*uintptr)(unsafe.Pointer(dst)) = *(*uintptr)(unsafe.Pointer(src))
}
func arg8(fun interface{}) (dst func(*uint8, *uint8) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint8, *uint8) int)
}
func arg32(fun interface{}) (dst func(*uint32, *uint32) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint32, *uint32) int)
}
func arg64(fun interface{}) (dst func(*uint64, *uint64) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint64, *uint64) int)
}
func u8(slice interface{}, size uintptr) (src []uint8) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint8)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func u32(slice interface{}, size uintptr) (src []uint32) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint32)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func u64(slice interface{}, size uintptr) (src []uint64) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint64)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
////////////////////////////////////////////////////////////////////////////////
func Top(compar, dst, src interface{}) {
	size := elemsize(src) //8,4,1

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		top64.Top(&m, arg64(compar), u64(dst, m[0]), u64(src, m[0]))
		return
	}

	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		top32.Top(&m, arg32(compar), u32(dst, m[0]), u32(src, m[0]))
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	top8.Top(&m, arg8(compar), u8(dst, m[0]), u8(src, m[0]))
	return
}
