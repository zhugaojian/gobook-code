package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a []int
	b := make([]int, 0)
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}

	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Printf("len=%d,cap=%d,type=%v\n", len(a), cap(a), as.Data)
	fmt.Printf("len=%d,cap=%d,type=%v\n", len(b), cap(b), bs.Data)
}
