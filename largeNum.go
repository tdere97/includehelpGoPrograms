package main

import "fmt"

func LargeNum() {
	var a, b, c int = 4, 6, 11

	fmt.Println("Largets number among :", a, b)
	if a > b {
		fmt.Println("is : ", a)
	} else {
		fmt.Println("is : ", b)
	}
	fmt.Println("Largets number among :", a, b, c)
	if a > b && a > c {
		fmt.Println("is : ", a)
	} else if b > a && b > c {
		fmt.Println("is : ", b)
	} else {
		fmt.Println("is : ", c)
	}
}
