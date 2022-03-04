package main

import "fmt"

func main() {

	fmt.Println("Slices Demo 2")

	slice := []int {1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	slice2 := slice[1:] // Copies from index 1, i.e 2, 3, 4, 5, 6
	fmt.Println(slice2) 

	slice3 := slice[1:3] // 2, 3
	fmt.Println(slice3)

	slice4 := slice[:4]  // Upto index 4, i.e 1, 2, 3, 4
	fmt.Println(slice4)

}