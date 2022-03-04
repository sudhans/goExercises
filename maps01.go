package main

import "fmt"

func main() {
	fmt.Println("Maps Demo")

	m := map[string]int{"foo": 42}

	fmt.Println(m)

	fmt.Println(m["foo"]) // Prints the value corresponds to the key , here 42

	m["foo"] = 27 // update value of a key
	fmt.Println(m)

	delete(m, "foo") // Delete a key in map.
	fmt.Println(m)   // Prints empty map in this case
}
