package main

import "fmt"

func main() {
	age := 32           // regular variable
	var agePointer *int // pointer
	agePointer = &age

	fmt.Println("Age pointer:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
