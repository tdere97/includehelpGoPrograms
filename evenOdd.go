package main

import "fmt"

func EvenOdd() {
	var a, b int = 2, 7

	if a%2 == 0 {
		fmt.Println("Number is even :", a)
	} else {
		fmt.Println("Number is odd :", a)
	}
	if b%2 == 0 {
		fmt.Println("Number is even :", b)
	} else {
		fmt.Println("Number is odd :", b)
	}
}
