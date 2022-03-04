package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "MSD"
	fmt.Printf("The name is %s \n", *firstName)

	fullName := "Master"
	fmt.Println(fullName)

	strPointer := &fullName
	fmt.Println(strPointer, *strPointer)
	

	fullName = "Sachin"
	fmt.Println(strPointer, *strPointer)
}
