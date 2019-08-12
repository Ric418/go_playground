package main

import (
	"./issues"
	"fmt"
)

func main() {
	if res := issues.Fib(10); res < 0{
		fmt.Printf("%d is minus value", res)
	} else{
		fmt.Print(res)
	}
}