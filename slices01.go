package main

import "fmt"

func main() {

	arr := [3]int{4, 5, 6}

	fmt.Println(arr)

	slice1 := arr[:] // slice points to the value the array is holding, kind of pointer
	fmt.Println(slice1)

	arr[1] = 7
	slice1[2] = 8

	fmt.Println(arr, slice1)

}
