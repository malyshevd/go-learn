package main

import "fmt"

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	prevRes := make(map[uint]uint)
	var n uint

nEnter:
	fmt.Print("Input positive number: ")
	fmt.Scanln(&n)

	res := prevRes[n]
	if res == 0 && n != 0 {
		res = fibonacci(n)
		prevRes[n] = res
	}

	fmt.Println("Fibonacci number: ", res)
	goto nEnter
}
