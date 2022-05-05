package main

var n int

func foo() (int, error) {
	return 1, nil
}

func g() {
	println(n)
}

func main() {
	n, _ := foo()

	g()

	println(n)
}
