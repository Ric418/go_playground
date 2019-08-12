package main

import (
	"fmt"
	"os"
	"strconv"

	//"strconv"
)

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fib(n int) (res int) {
	switch {
	case n < 0:
		return n
	default:
		f := fibonacci()
		for i := 0; i < n-1; i++ {
			f()
		}
		return f()
	}
}

func main() {
	if len(os.Args) < 3 {
		var i int
		n := os.Args[1]
		i, _ = strconv.Atoi(n)
		if res := fib(i); res < 0 {
			fmt.Printf("%d is minus value\n", res)
			os.Exit(1)
		} else {
			fmt.Printf("fibnacci number: %d \n", res)
			os.Exit(1)
		}
	} else {
		fmt.Printf("error: Too many args.\n")
		os.Exit(1)
	}
}

