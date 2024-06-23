package main

import "fmt"

func Swap2Num() {

	a, b := 5, 6
	fmt.Println("Before", a, b)
	a, b = b, a
	fmt.Println("After", a, b)
}
