package main

import "fmt"

const (
	first  = iota + 6
	second = 2 << iota // value of iota will be 1 here, left shift 2 by one bit which is equivalent to multiply 2 by 2 and the result is 4
)

const (
	third = iota // value of iota is reset to 0 as this is a new block
)

func main() {
	const pi = 3.1415
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)
	fmt.Println(c + 1.2) // This will print 4.2 as c's type is not specified and only inferred

	const d int = 4
	var f float32 = 3.14
	//fmt.Println(d + f) // Untyped float constant truncated to int
	fmt.Println(float32(d) + f) // Typecast d to float32 to enable addition with f

	fmt.Println(first, second, third) // prints 6 4 0

}
