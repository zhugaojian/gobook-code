package main

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t:=5
	defer func() {
		t=t+5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r=r+5
	}(r)
	return 1
}

func main() {
	println("f1=",f1())
	println("f2=",f2())
	println("f3=",f3())
}