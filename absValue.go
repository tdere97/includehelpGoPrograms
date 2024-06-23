package main

import (
	"fmt"
	"math"
)

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func AbsValues() {
	var f float64 = -12.34
	var num int = -10
	fmt.Println("Absolute value of float is : ", math.Abs(f))
	fmt.Println("Absolute value of integer is : ", absInt(num))

}
