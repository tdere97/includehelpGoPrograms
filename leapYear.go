package main

import "fmt"

/*
To check if a year is a leap year, you can use the following conditions:
A year is a leap year if it is divisible by 4.
However, if the year is divisible by 100, it is not a leap year, unless:
The year is also divisible by 400.
*/
func LeapYear() {
	var year int
	fmt.Println("Enter a year")
	//fmt.Scanf("%d", &year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Year is leaf year : ", year)
	} else {
		fmt.Println("Year is not a leaf year : ", year)
	}
}
