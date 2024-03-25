package main

import "fmt"

func main() {
	age := 32           // regular variable
	var agePointer *int // pointer
	agePointer = &age
	//fmt.Println("Age:", age)
	fmt.Println("Age pointer:", *agePointer)
	// adultYears := getAdultYears(age)
	// fmt.Println("Adult Years:", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
