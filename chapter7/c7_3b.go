package main

func f4() int {
	r := 0
	defer func() {
		r++
	}()
	return r
}

func f5() int {
	r := 0
	defer func(i int) {
		i++
	}(r)
	return 0
}

func main() {
	println("f4=", f4())
	println("f5=", f5())
}
