package main

func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a,a)
		a=a+i
		return a
	}
}

func main() {
	f:=fa(1)
	println(f(1))
	println(f(2))

}