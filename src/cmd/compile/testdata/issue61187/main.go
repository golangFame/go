package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var slice = []byte{'H', 'e', 'l', 'l', 'o', ','}

func main() {
	ptr := uintptr(unsafe.Pointer(&slice)) + 100
	header := (*reflect.SliceHeader)(unsafe.Pointer(ptr))
	header.Data += 1 // comment this line does not lead to internal compiler error
	fmt.Printf("%d %d\n", cap(slice), header.Cap)
}
