package main

import "fmt"

// c= (f-32)* 5/9
// f= (c *9/5)+32
func FarTOCel() {
	var temp float64
	fmt.Println("Enter temp in Fahrenheit")
	//fmt.Scanf("%f", &temp)
	fmt.Println("Temp in Fahrenheit : ", temp)
	ctemp := (temp - 32) / 1.8
	fmt.Println("Temp in Celsius : ", ctemp)

	fmt.Println("Enter temp in Celsius")
	//fmt.Scanf("%f", &temp)
	fmt.Println("Temp in Celsius : ", temp)
	temp = (temp * 1.8) + 32
	fmt.Println("Temp in Fahrenheit : ", temp)
}
