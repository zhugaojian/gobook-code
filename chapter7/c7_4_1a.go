package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1: 2, 2: 3}
	var d [3]int
	fmt.Printf("len=%d,value=%v\n", len(a), a)
	fmt.Printf("len=%d,value=%v\n", len(b), b)
	fmt.Printf("len=%d,value=%v\n", len(c), c)
	fmt.Printf("len=%d,value=%v\n", len(d), d)
}
