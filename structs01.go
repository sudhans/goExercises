package main

import "fmt"

func main() {
	fmt.Println("Structs Demo")

	type employee struct {
		ID int
		FirstName string
		LastName string
	}

	var msd employee
	fmt.Println(msd)

	msd.ID = 125946
	msd.FirstName = "Madhusudhan"
	msd.LastName = "SD"

	fmt.Println(msd)

	msd2 := employee{ID:125946, FirstName: "Madhu", LastName: "SD"}
    fmt.Println(msd2)

}
