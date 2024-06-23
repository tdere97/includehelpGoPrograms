package main

import "fmt"

func BitwiseOperator() {
	var num1, num2 int = 6, 2
	fmt.Printf("%d in Binary is %b\n", num1, num1)
	fmt.Printf("%d in Binary is %b\n", num2, num2)
	fmt.Printf("Bitwise & of \n%b\n%b\n%b\n", num1, num2, num1&num2)
	fmt.Printf("Bitwise | of \n%b\n%b\n%b\n", num1, num2, num1|num2)
	fmt.Printf("Bitwise ^ of \n%b\n%b\n%b\n", num1, num2, num1^num2)
}
