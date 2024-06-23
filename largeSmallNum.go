package main

import (
	"fmt"
	"math"
)

func LargeSmallNum() {
	var a, b float64 = 10, 34

	fmt.Println("Maxnumber is :", math.Max(a, b))
	fmt.Println("Minnumber is :", math.Min(a, b))

}
