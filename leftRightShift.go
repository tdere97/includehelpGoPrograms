package main

import "fmt"

func LeftRightShift() {
	var num int = 4
	fmt.Printf("%d in Binary is %b\n", num, num)

	num = num << 3
	fmt.Println("Num left shifted by 3")
	fmt.Printf("%d in Binary is %b\n", num, num)

	num = num >> 3
	fmt.Println("Num right shifted by 3")
	fmt.Printf("%d in Binary is %b\n", num, num)

}
